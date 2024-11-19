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
    return <div>Settings for {this.widgetId}</div>;
  }
}
