# Streaming Widget

The Streaming Widget is a key feature that allows users to display their stats in real-time (e.g., on OBS).

## Overview
The widget renders a set of cards:
- **Rating Overview:** Summary of rating battles.
- **Unrated Overview:** Summary of regular battles (WN8, Winrate, etc.).
- **Vehicle Cards:** Individual stats for recently played vehicles.

## Configuration & Preview
Users can configure the widget via:
1.  **Public Preview:** `/widget` (Mock data) and `/widget/account/:id` (Real data).
2.  **Custom Widget Editor:** `/app/widgets/new` (Authenticated).

### Layout & UX
The preview pages use a **Split-Pane Grid Layout** on desktop:
- **Left Panel (Settings):** A fixed-width sidebar containing configuration controls and player search. This ensures UI elements remain stable and don't "jump" when the preview updates.
- **Right Panel (Preview):** An **OBS Mockup** window that fills available vertical space. The widget content inside is scrollable, while the OBS chrome (toolbar, bottom panels) stays fixed.

On mobile devices, only the settings panel is shown (no preview). Users configure their widget on mobile and copy the OBS link to use on desktop.

### Client-Side Preview
To ensure a smooth user experience, the preview pages utilize a **Client-Side Toggle** approach:
- **Server-Side Rendering (SSR):** The server renders *all* potential components (e.g., both overviews and up to 10 vehicles) but hides the ones not currently selected using CSS (`hidden` class).
- **Interactive Toggling:** A lightweight JavaScript function (`updateWidgetPreview`) listens to changes in the settings form (Checkboxes, Slider).
- **Instant Feedback:** When a setting changes, the script toggles the visibility classes immediately without making a network request.
- **URL Persistence:** The script updates the browser's URL (via `pushState`) so configurations can be shared or copied (e.g., "Copy OBS Link").

## Key Files
- `cmd/frontend/components/widget/default.templ`: Defines the widget HTML structure, rendering all components with appropriate ID/data attributes.
- `cmd/frontend/components/obs.templ`: The OBS Mockup component (full-height with scrollable content area).
- `cmd/frontend/routes/widget/index.templ`: Public mock preview page + Client-side toggle script.
- `cmd/frontend/routes/widget/preview.templ`: Account specific preview page + Client-side toggle script.
- `cmd/frontend/components/widget/settings.templ`: Reusable settings form component.

## Architecture Decisions
- **Why Client-Side?** The backend cost to fetch stats for 1 vehicle vs 10 vehicles is negligible. By sending all data upfront, we eliminate "flash of unstyled content" (FOUC) and UI freezing associated with repeated network requests for minor layout changes.
- **Why Grid Layout?** A `flex-wrap` layout caused the settings panel to resize and jump whenever the widget (preview) changed size. A CSS Grid with a fixed sidebar prevents this layout shifting, providing a professional and stable feel.
- **HTMX:** Used elsewhere in the app, but for the *preview* specifically, vanilla JS class toggling was chosen for maximum speed and simplicity.

