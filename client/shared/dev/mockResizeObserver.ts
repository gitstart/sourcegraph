import { vi } from 'vitest';

import ResizeObserver from 'resize-observer-polyfill'

if ('ResizeObserver' in window === false) {
    window.ResizeObserver = ResizeObserver
}

vi.mock('use-resize-observer', () => ({
    __esModule: true,
    // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
    default: await vi.importActual('use-resize-observer/polyfilled'),
}))
