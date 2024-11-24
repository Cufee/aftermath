import { Component, Method, Prop, State, h } from '@stencil/core';
import Cards from './Cards';

@Component({
  tag: 'amth-widget',
  styleUrl: 'widget.css',
  shadow: true,
})
export class Widget {
  /**
   * Custom widget ID
   */
  @Prop({ attribute: 'widget-id', mutable: true }) widgetId: string;
  /**
   * Wargaming account ID
   */
  @Prop({ attribute: 'account-id', mutable: true }) accountId: string;
  /**
   * If set to true, the widget will automatically fetch data when needed
   */
  @Prop({ attribute: 'auto-reload', mutable: true }) autoReload: boolean = false;
  /**
   * Overwrite Aftermath WebSocket host
   */
  @Prop({ attribute: 'ws-overwrite' }) wsHost: string = 'wss://amth.one';
  /**
   * Overwrite Aftermath API host for all requests
   */
  @Prop({ attribute: 'api-overwrite' }) apiHost: string = 'https://amth.one';
  /**
   * Data to initialize the widget with
   */
  @Prop({ attribute: 'initial-data' }) initialData?: string;

  @State() loadingSpinner: boolean = true;
  @State() fetchInProgress: boolean = false;

  @State() error: string | null = null;
  @State() widgetCards: Cards | null = null;
  @State() widgetOptions: Options | null = null;
  @State() widgetAccount: Account | null = null;
  @State() websocket: WebSocket | null = null;
  @State() autoReloadTimer: NodeJS.Timer | null = null;

  componentWillLoad() {
    this.setAutoReload(this.autoReload);
    this.initWebSocketConnection();

    if (this.initialData) {
      const data = this.parseWidgetData(this.initialData);
      if (data) {
        this.accountId = data.account.id;
        this.widgetCards = data.cards;
        this.widgetAccount = data.account;
        this.widgetOptions = data.options;
        this.loadingSpinner = false;
        return;
      }
    }
    this.refresh();
  }

  @Method()
  async refresh() {
    this.loadingSpinner = true;
    return this.fetchWidgetData()
      .then(res => {
        if (res.ok !== true) {
          this.error = res.error;
        } else {
          this.accountId = res.data.account.id;
          this.widgetCards = res.data.cards;
          this.widgetAccount = res.data.account;
          this.widgetOptions = this.widgetOptions || res.data.options;
        }
      })
      .finally(() => (this.loadingSpinner = false));
  }

  @Method()
  async options(): Promise<Options | null> {
    return this.widgetOptions ?? null;
  }

  @Method()
  async setOptions(opts: Options) {
    console.debug('updated settings');
    this.widgetOptions = opts;
  }

  // removes widget id, replaces account id, does not change options
  @Method()
  async setAccountId(id: string) {
    this.loadingSpinner = true;
    this.widgetAccount = null;
    this.widgetCards = null;
    this.accountId = id;
    this.widgetId = '';
    this.refresh();
  }

  // removes account id, replaces widget id, removes options
  @Method()
  async setWidgetId(id: string) {
    this.loadingSpinner = true;
    this.widgetOptions = null;
    this.widgetAccount = null;
    this.widgetCards = null;
    this.accountId = '';
    this.widgetId = id;
    this.refresh();
  }

  @Method()
  async setAutoReload(value: boolean) {
    this.autoReload = value === true;
    if (this.autoReloadTimer) clearInterval(this.autoReloadTimer);

    if (this.autoReload) {
      console.debug('enabled auto reload');
      this.autoReloadTimer = setInterval(this.fetchDataOnBattle, 10000);
    } else {
      console.debug('disabled auto reload');
    }
  }

  private initWebSocketConnection() {
    if (!this.widgetId || this.websocket) return;

    this.websocket = new WebSocket(`${this.wsHost}/api/p/realtime/widget/custom/${this.widgetId}/`);
    this.websocket.addEventListener('open', () => {
      console.debug('connected to realtime api');
    });
    this.websocket.addEventListener('close', () => {
      console.debug('disconnected, will attempt to reconnect in 30 seconds');
      setTimeout(this.initWebSocketConnection, 30000);
      this.websocket = null;
    });
    this.websocket.addEventListener('error', (event: Event & { data: string }) => {
      console.error('websocket connection error, will attempt to reconnect in 30 seconds', event.data);
      setTimeout(this.initWebSocketConnection, 30000);
      try {
        this.websocket?.close();
      } catch (error) {
        console.error(error);
      }
      this.websocket = null;
    });
    this.websocket.addEventListener('message', async event => {
      let data: { command: string } | null = null;
      try {
        data = JSON.parse(event.data);
        console.debug('realtime api message', data);
      } catch (error) {
        console.error('invalid message received on websocket', error);
        return;
      }

      if (data.command === 'reload') {
        console.debug('realtime api requested a reload');

        const result = await this.fetchWidgetData();
        if (result.ok !== true) {
          console.error('failed to fetch widget data', result.error);
          return;
        }
        console.log(result.data);
        this.accountId = result.data.account.id;
        this.widgetCards = result.data.cards;
        this.widgetAccount = result.data.account;
        this.widgetOptions = result.data.options;
      }
    });
  }

  private parseWidgetData(input: string): WidgetData | null {
    try {
      const data = JSON.parse(input);
      console.debug('loaded initial widget data');
      return data;
    } catch (error) {
      console.error('failed to parse initial-data', error);
      return null;
    }
  }

  private async fetchWidgetData(): Promise<{ ok: true; data: WidgetData } | { ok: false; error: string }> {
    if (this.fetchInProgress) return;
    this.fetchInProgress = true;

    if (!this.accountId && !this.widgetId) {
      return { ok: false, error: 'account-id or widget-id parameter is required' };
    }

    const endpoint = !this.widgetId ? `/widget/account/${this.accountId}/live/json/` : `/widget/custom/${this.widgetId}/live/json/`;
    try {
      const url = this.apiHost + endpoint;
      console.debug(url);
      const res = await fetch(url);
      const data = await res.json();
      if (data.error) {
        console.error('widget api returned an error', data.error);
        return { ok: false, error: data.error };
      }
      return { ok: true, data };
    } catch (error) {
      console.error(error);
      return { ok: false, error: error.toString() };
    } finally {
      this.fetchInProgress = false;
    }
  }

  private fetchDataOnBattle() {
    if (!!this.error || !this.widgetAccount || this.fetchInProgress) return;

    const lastBattleTime = Date.parse(this.widgetAccount.lastBattleTime);
    if (!lastBattleTime || Number.isNaN(lastBattleTime)) return;

    let apiHost = '';
    switch (this.widgetAccount.realm) {
      case 'na':
        apiHost = 'wotblitz.com';
        break;
      case 'eu':
        apiHost = 'wotblitz.eu';
        break;
      case 'as':
        apiHost = 'wotblitz.asia';
        break;
      default:
        console.error('invalid account realm', this.widgetAccount.realm);
        return;
    }

    fetch(`https://api.${apiHost}/wotb/account/info/?application_id=f44aa6f863c9327c63ba26be3db0d07f&account_id=${this.widgetAccount.id}&fields=last_battle_time`)
      .then(response => response.json())
      .then(async data => {
        if (data.data[this.widgetAccount.id.toString()].last_battle_time <= lastBattleTime) {
          console.debug('no new battles since last refresh');
        }
        console.debug('found a new battle since last refresh, fetching new data');

        const result = await this.fetchWidgetData();
        if (result.ok !== true) {
          console.error('failed to fetch widget data', result.error);
          return;
        }
        this.widgetCards = result.data.cards;
        this.widgetAccount = result.data.account;
        this.widgetOptions = this.widgetOptions || result.data.options;
      })
      .catch(error => {
        console.error('request to wargaming api failed', error);
      });
  }

  render() {
    if (this.loadingSpinner === true) {
      return (
        <div class="w-full flex flex-col items-center justify-center gap-2">
          <div class="size-8">
            <img src={this.apiHost + '/assets/icon/64.png'} />
          </div>
          <span class="text-gray-400 text-sm"> Loading Widget </span>
          <div class="grid w-full place-items-center overflow-x-scroll rounded-lg lg:overflow-visible">
            <svg class="text-gray-400 animate-spin" viewBox="0 0 64 64" fill="none" xmlns="http://www.w3.org/2000/svg" width="24" height="24">
              <path
                d="M32 3C35.8083 3 39.5794 3.75011 43.0978 5.20749C46.6163 6.66488 49.8132 8.80101 52.5061 11.4939C55.199 14.1868 57.3351 17.3837 58.7925 20.9022C60.2499 24.4206 61 28.1917 61 32C61 35.8083 60.2499 39.5794 58.7925 43.0978C57.3351 46.6163 55.199 49.8132 52.5061 52.5061C49.8132 55.199 46.6163 57.3351 43.0978 58.7925C39.5794 60.2499 35.8083 61 32 61C28.1917 61 24.4206 60.2499 20.9022 58.7925C17.3837 57.3351 14.1868 55.199 11.4939 52.5061C8.801 49.8132 6.66487 46.6163 5.20749 43.0978C3.7501 39.5794 3 35.8083 3 32C3 28.1917 3.75011 24.4206 5.2075 20.9022C6.66489 17.3837 8.80101 14.1868 11.4939 11.4939C14.1868 8.80099 17.3838 6.66487 20.9022 5.20749C24.4206 3.7501 28.1917 3 32 3L32 3Z"
                stroke="currentColor"
                stroke-width="5"
                stroke-linecap="round"
                stroke-linejoin="round"
              ></path>
              <path
                d="M32 3C36.5778 3 41.0906 4.08374 45.1692 6.16256C49.2477 8.24138 52.7762 11.2562 55.466 14.9605C58.1558 18.6647 59.9304 22.9531 60.6448 27.4748C61.3591 31.9965 60.9928 36.6232 59.5759 40.9762"
                stroke="rgb(209 213 219)"
                stroke-width="5"
                stroke-linecap="round"
                stroke-linejoin="round"
                class="text-gray-900"
              ></path>
            </svg>
          </div>
        </div>
      );
    }
    if (this.error) {
      return (
        <div class="w-full flex flex-col items-center justify-center">
          <svg class="text-red-300 size-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126ZM12 15.75h.007v.008H12v-.008Z"
            />
          </svg>
          <span class="text-gray-400"> {this.error} </span>
        </div>
      );
    }

    return <Cards cards={this.widgetCards} options={this.widgetOptions} assetsDomain={this.apiHost} />;
  }
}
