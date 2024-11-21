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
  @Prop({ attribute: 'ws-overwrite' }) wsHostOverwrite?: string;
  /**
   * Overwrite Aftermath API host for all requests
   */
  @Prop({ attribute: 'api-overwrite' }) apiHostOverwrite?: string;
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

  private defaultApiDomain = 'https://amth.one';
  private defaultWsDomain = 'wss://amth.one';

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

    this.websocket = new WebSocket(`${this.wsHostOverwrite ?? this.defaultWsDomain}/api/p/realtime/widget/custom/${this.widgetId}/`);
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

    const apiDomain = this.apiHostOverwrite ?? this.defaultApiDomain;
    const endpoint = !!this.accountId ? `/widget/account/${this.accountId}/live/json/` : `/widget/custom/${this.widgetId}/live/json/`;
    try {
      const url = apiDomain + endpoint;
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
      return <div>loading</div>;
    }
    if (this.error) {
      return <div>error: {this.error}</div>;
    }

    return <Cards cards={this.widgetCards} options={this.widgetOptions} />;
  }
}
