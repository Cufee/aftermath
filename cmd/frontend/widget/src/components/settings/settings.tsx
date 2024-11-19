import { Component, Prop, State, Watch, h } from '@stencil/core';
import { Widget } from '../widget/widget';

@Component({
  tag: 'amth-widget-settings',
  styleUrl: 'settings.css',
  shadow: true,
})
export class Settings {
  /**
   * Widget ID this instance of settings is controlling
   */
  @Prop({ attribute: 'widget-id' }) widgetId: string;

  @State() error: string | null = null;
  @State() loadingState: boolean = true;
  @State() widgetOptions: Options | null = null;
  @State() widget: (HTMLElement & Widget) | null = null;

  @Watch('widgetOptions')
  watchPropHandler(newValue: Options, oldValue: Options) {
    this.handleOptionChange(newValue, oldValue);
  }

  async componentWillLoad() {
    try {
      const target = document.getElementById(this.widgetId);
      if (!target?.['options']) {
        this.error = 'Widget not found';
        this.loadingState = false;
        return;
      }
      this.widget = target as any;
      const opts = await this.widget.options();
      this.widgetOptions = opts;
    } catch (error) {
      this.error = error.toString();
    } finally {
      this.loadingState = false;
    }
  }

  private handleOptionChange(newValue: Options, oldValue?: Options) {
    oldValue = oldValue || this.widgetOptions;
    if (!newValue || !this.widget) return;
    this.saveChanges();
  }

  private toggleCardVisible(card: 'rating' | 'unrated' | 'vehicles') {
    const newValue: Options = JSON.parse(JSON.stringify(this.widgetOptions));
    switch (card) {
      case 'rating':
        newValue.rating.visible = !newValue.rating.visible;
        break;
      case 'unrated':
        newValue.unrated.visible = !newValue.unrated.visible;
        break;
      case 'vehicles':
        newValue.vehicles.visible = !newValue.vehicles.visible;
        break;
    }
    this.widgetOptions = newValue;
  }

  private setVehicleLimit(limit: number) {
    if (limit < 0 || limit > 10) return;
    const newValue: Options = JSON.parse(JSON.stringify(this.widgetOptions));
    newValue.vehicles.limit = limit;
    this.widgetOptions = newValue;
  }

  private saveChanges() {
    this.widget.setOptions(this.widgetOptions);
  }

  render() {
    if (this.loadingState) {
      return <div>Loading widget settings</div>;
    }
    if (this.error) {
      return <div>Error: {this.error}</div>;
    }

    return (
      <div class="form-control flex flex-col gap-2">
        <div class="flex flex-col p-4">
          <span class="text-lg">Rating Battles</span>
          <label class="label group">
            <span class="label-text group-hover:underline">Show Overview Card</span>
            <input
              type="checkbox"
              checked={this.widgetOptions.rating.visible}
              onChange={() => this.toggleCardVisible('rating')}
              class="toggle toggle-secondary transition-all duration-250 ease-in-out"
            />
          </label>
        </div>
        <div class="flex flex-col p-4">
          <span class="text-lg">Regular Battles</span>
          <label class="label group">
            <span class="label-text group-hover:underline">Show Overview Card</span>
            <input
              type="checkbox"
              checked={this.widgetOptions.unrated.visible}
              onChange={() => this.toggleCardVisible('unrated')}
              class="toggle toggle-secondary transition-all duration-250 ease-in-out"
            />
          </label>
          <label class="label flex flex-col items-start gap-1 group">
            <span class="label-text group-hover:underline">Vehicle Cards</span>
            <input
              id="widget-settings-vl"
              type="range"
              min="0"
              max="10"
              class="range w-full"
              step="1"
              value={this.widgetOptions.vehicles.limit}
              onChange={e => this.setVehicleLimit(parseInt(e.target?.['value']) ?? 10)}
            />
            <div class="flex w-full justify-between px-2 text-xs">
              {[...Array(11)].map((_, i) => (
                <div class="flex flex-col items-center">
                  <span>{i}</span>
                </div>
              ))}
            </div>
          </label>
        </div>
      </div>
    );
  }
}
