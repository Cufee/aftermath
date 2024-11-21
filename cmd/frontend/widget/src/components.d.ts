/* eslint-disable */
/* tslint:disable */
/**
 * This is an autogenerated file created by the Stencil compiler.
 * It contains typing information for all components that exist in this project.
 */
import { HTMLStencilElement, JSXBase } from "@stencil/core/internal";
export namespace Components {
    interface AmthWidget {
        /**
          * Wargaming account ID
         */
        "accountId": string;
        /**
          * Overwrite Aftermath API host for all requests
         */
        "apiHost": string;
        /**
          * If set to true, the widget will automatically fetch data when needed
         */
        "autoReload": boolean;
        /**
          * Data to initialize the widget with
         */
        "initialData"?: string;
        "options": () => Promise<Options | null>;
        "refresh": () => Promise<void>;
        "setAccountId": (id: string) => Promise<void>;
        "setAutoReload": (value: boolean) => Promise<void>;
        "setOptions": (opts: Options) => Promise<void>;
        "setWidgetId": (id: string) => Promise<void>;
        /**
          * Custom widget ID
         */
        "widgetId": string;
        /**
          * Overwrite Aftermath WebSocket host
         */
        "wsHost": string;
    }
    interface AmthWidgetSettings {
        /**
          * Widget ID this instance of settings is controlling
         */
        "widgetId": string;
    }
}
declare global {
    interface HTMLAmthWidgetElement extends Components.AmthWidget, HTMLStencilElement {
    }
    var HTMLAmthWidgetElement: {
        prototype: HTMLAmthWidgetElement;
        new (): HTMLAmthWidgetElement;
    };
    interface HTMLAmthWidgetSettingsElement extends Components.AmthWidgetSettings, HTMLStencilElement {
    }
    var HTMLAmthWidgetSettingsElement: {
        prototype: HTMLAmthWidgetSettingsElement;
        new (): HTMLAmthWidgetSettingsElement;
    };
    interface HTMLElementTagNameMap {
        "amth-widget": HTMLAmthWidgetElement;
        "amth-widget-settings": HTMLAmthWidgetSettingsElement;
    }
}
declare namespace LocalJSX {
    interface AmthWidget {
        /**
          * Wargaming account ID
         */
        "accountId"?: string;
        /**
          * Overwrite Aftermath API host for all requests
         */
        "apiHost"?: string;
        /**
          * If set to true, the widget will automatically fetch data when needed
         */
        "autoReload"?: boolean;
        /**
          * Data to initialize the widget with
         */
        "initialData"?: string;
        /**
          * Custom widget ID
         */
        "widgetId"?: string;
        /**
          * Overwrite Aftermath WebSocket host
         */
        "wsHost"?: string;
    }
    interface AmthWidgetSettings {
        /**
          * Widget ID this instance of settings is controlling
         */
        "widgetId"?: string;
    }
    interface IntrinsicElements {
        "amth-widget": AmthWidget;
        "amth-widget-settings": AmthWidgetSettings;
    }
}
export { LocalJSX as JSX };
declare module "@stencil/core" {
    export namespace JSX {
        interface IntrinsicElements {
            "amth-widget": LocalJSX.AmthWidget & JSXBase.HTMLAttributes<HTMLAmthWidgetElement>;
            "amth-widget-settings": LocalJSX.AmthWidgetSettings & JSXBase.HTMLAttributes<HTMLAmthWidgetSettingsElement>;
        }
    }
}
