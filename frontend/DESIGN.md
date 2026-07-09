# Design System: Neobrutalist Career Catalyst

## Overview
This design system is based on the Stitch-generated "Resume Engine" landing page with soft pastel neobrutalist aesthetics.

## Color Palette

| Token | Hex | Usage |
|-------|-----|-------|
| Primary | `#ffe600` | CTAs, highlights, score gauges, hero accent |
| Secondary | `#ff00f5` | Hover states, accent tags, secondary elements |
| Tertiary | `#00f0ff` | Links, info tags, technical data |
| Black | `#000000` | Borders, shadows, text, heavy elements |
| White | `#ffffff` | Card backgrounds |
| Surface | `#f9f9f9` | Page background |
| Card | `#ffffff` | Card backgrounds (same as white) |
| Text | `#1b1b1b` | Primary text |
| Text Muted | `#4b4731` | Secondary text, descriptions |

## Typography

### Font Families
- **Headlines**: `Space Grotesk` (self-hosted, bold weight)
- **Body**: `Work Sans` (self-hosted, various weights)
- **Fallback**: `system-ui, sans-serif`

### Scale
- Display/hero: 72px (7xl) - bold uppercase
- Headings: 32-48px (3xl-4xl) - bold uppercase
- Body: 16-20px (base-lg) - regular/medium
- Labels: 14px - bold uppercase

### Letter Spacing
- Headlines: `-0.02em` (tight)
- Labels: `tracking-wide` or `uppercase`

## Spacing System

- Base unit: 4px
- xs: 8px | sm: 16px | md: 32px | lg: 64px | xl: 128px
- Section padding: 64px-96px (py-16 to py-24)
- Card padding: 32px (p-8)
- Component gap: 24px (gap-6)

## Borders & Shadows

- **Border width**: 4px solid black (primary elements)
- **Border radius**: 0px (sharp corners - neobrutalist)
- **Shadows**: Hard offset, no blur. Single `4px 4px 0 0 #000` used on every shadowed element.

## Components

### Buttons

**Primary (btn-neo)**:
- Background: primary (#FFE600)
- Text: black, bold, uppercase
- Border: 4px solid black
- Shadow: 4px 4px 0 0 #000
- Pressed (hover/active): translate(4px, 4px), shadow 0 0 0 0

**Secondary**:
- Background: white
- Text: black, bold, uppercase
- Border: 4px solid black
- Shadow: 4px 4px 0 0 #000

### Cards (card-neo)
- Background: white (#ffffff)
- Border: 4px solid black
- Shadow: 4px 4px 0 0 #000
- Padding: 32px (p-8)

### Tags (tag-neo)
- Border: 2px solid black
- Text: black, bold, uppercase
- Padding: 8px 16px (px-4 py-2)

### Inputs
- Border: 4px solid black
- Focus: ring with tertiary color

## Icon Style
- Stroke-based icons from lucide-react
- Stroke width: 2px
- Color: black (fill for accent icons)

## Responsive Breakpoints
- Mobile: < 768px (stack vertically)
- Tablet: 768px - 1024px (hybrid)
- Desktop: > 1024px (full horizontal layouts)

## Implementation Notes

1. All components use Tailwind CSS utilities
2. Custom utilities defined with `@utility` and component classes in `@layer components` in index.css
3. Self-hosted fonts in public/fonts/
4. No inline styles - all utility classes
5. Brutalist toggle: instant (no animation)
