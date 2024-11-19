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
  @State() data: WidgetData | null = null;
  @State() websocket: WebSocket | null = null;
  @State() autoReloadTimer: NodeJS.Timer | null = null;

  private defaultApiDomain = 'https://amth.one';
  private defaultWsDomain = 'wss://amth.one';

  componentWillLoad() {
    this.setAutoReload(this.autoReload);

    if (this.initialData) {
      this.data = this.parseWidgetData(this.initialData);
      if (this.data) {
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
          this.data = res.data;
        }
      })
      .finally(() => (this.loadingSpinner = false));
  }

  @Method()
  async setAccountId(id: string) {
    this.loadingSpinner = true;
    this.accountId = id;
    this.widgetId = '';
    this.data = null;
    this.refresh();
  }

  @Method()
  async setWidgetId(id: string) {
    this.loadingSpinner = true;
    this.accountId = '';
    this.widgetId = id;
    this.data = null;
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
      } catch (error) {
        console.error('invalid message received on websocket', error);
      }

      if (data.command === 'reload') {
        const result = await this.fetchWidgetData();
        if (result.ok !== true) {
          console.error('failed to fetch widget data', result.error);
          return;
        }
        this.data = result.data;
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

    const accountId = this.data?.account.id || this.accountId;
    if (!accountId && !this.widgetId) {
      return { ok: false, error: 'account-id or widget-id parameter is required' };
    }

    const apiDomain = this.apiHostOverwrite ?? this.defaultApiDomain;
    const endpoint = !!accountId ? `/widget/account/${accountId}/live/json/` : `/widget/custom/${this.widgetId}/live/json/`;
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
    if (!!this.error || !this.data || !this.data.account.id || !this.data.account.realm || this.fetchInProgress) return;

    const lastBattleTime = Date.parse(this.data.account.lastBattleTime);
    if (!lastBattleTime || Number.isNaN(lastBattleTime)) return;

    let apiHost = '';
    switch (this.data.account.realm) {
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
        console.error('invalid account realm', this.data.account.realm);
        return;
    }

    fetch(`https://api.${apiHost}/wotb/account/info/?application_id=f44aa6f863c9327c63ba26be3db0d07f&account_id=${this.data.account.id}&fields=last_battle_time`)
      .then(response => response.json())
      .then(async data => {
        if (data.data[this.data.account.id.toString()].last_battle_time <= lastBattleTime) {
          console.debug('no new battles since last refresh');
        }
        console.debug('found a new battle since last refresh, fetching new data');

        const result = await this.fetchWidgetData();
        if (result.ok !== true) {
          console.error('failed to fetch widget data', result.error);
          return;
        }
        this.data = result.data;
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

    return <Cards data={this.data} />;
  }
}
