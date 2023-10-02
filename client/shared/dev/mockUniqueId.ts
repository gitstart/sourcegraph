// The value returned by the `_.uniqueId()` function depends on the number of previous calls
// because it maintains an internal variable that increments on each call.
// Mock `_.uniqueId()` to avoid updating snapshots once the sequence or number of tests using this helper changes.
vi.mock('lodash', () =>
    Object.assign(await vi.importActual('lodash'), {
        // Keep all actual implementations except uniqueId
        uniqueId: (prefix = '') => `${prefix}test-id`,
    })
)

// The module factory of `jest.mock()` is not allowed to reference any out-of-scope variables

vi.mock(
    'lodash/uniqueId',
    () =>
        (prefix = '') =>
            `${prefix}test-id`
)
