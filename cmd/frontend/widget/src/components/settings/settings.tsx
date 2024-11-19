import { Component, Prop, h } from '@stencil/core';

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

  render() {
    const widget = document.getElementById(this.widgetId);
    if (!widget) {
      return <div>Widget not found</div>;
    }

    return <div>Settings for {this.widgetId}</div>;
  }
}
