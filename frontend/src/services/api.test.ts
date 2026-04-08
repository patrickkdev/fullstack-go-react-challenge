import axios from 'axios'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

// Mock localStorage
const localStorageMock = {
	getItem: vi.fn(),
	setItem: vi.fn(),
	removeItem: vi.fn(),
	clear: vi.fn(),
}
Object.defineProperty(window, 'localStorage', {
	value: localStorageMock,
})

vi.mock('axios')
const mockedAxios = vi.mocked(axios)

describe('api', () => {
	beforeEach(() => {
		vi.clearAllMocks()
		localStorageMock.getItem.mockClear()
		localStorageMock.removeItem.mockClear()
	})

	afterEach(() => {
		vi.restoreAllMocks()
	})

	describe('request interceptor', () => {
		it('adds authorization header when token exists', async () => {
			const token = 'test-token'
			localStorageMock.getItem.mockReturnValue(token)

			mockedAxios.create.mockReturnValue({
				...mockedAxios,
				interceptors: {
					request: {
						use: vi.fn((fn) => {
							const config = { headers: {} }
							fn(config)
							expect(config.headers.Authorization).toBe(`Bearer ${token}`)
						}),
					},
					response: { use: vi.fn() },
				},
			} as any)

			// Re-import to trigger the interceptor setup
			await import('./api')
		})

		it('does not add authorization header when no token', async () => {
			localStorageMock.getItem.mockReturnValue(null)

			mockedAxios.create.mockReturnValue({
				...mockedAxios,
				interceptors: {
					request: {
						use: vi.fn((fn) => {
							const config = { headers: {} }
							fn(config)
							expect(config.headers.Authorization).toBeUndefined()
						}),
					},
					response: { use: vi.fn() },
				},
			} as any)

			await import('./api')
		})
	})

	describe('response interceptor', () => {
		it('removes token on 401 error for non-auth endpoints', async () => {
			const error = {
				response: { status: 401 },
				config: { url: '/api/protected' },
			}

			mockedAxios.create.mockReturnValue({
				...mockedAxios,
				interceptors: {
					request: { use: vi.fn() },
					response: {
						use: vi.fn((_, errorHandler) => {
							errorHandler(error)
							expect(localStorageMock.removeItem).toHaveBeenCalledWith('authToken')
						}),
					},
				},
			} as any)

			await import('./api')
		})

		it('does not remove token on 401 for login endpoint', async () => {
			const error = {
				response: { status: 401 },
				config: { url: '/api/auth/login' },
			}

			mockedAxios.create.mockReturnValue({
				...mockedAxios,
				interceptors: {
					request: { use: vi.fn() },
					response: {
						use: vi.fn((_, errorHandler) => {
							errorHandler(error)
							expect(localStorageMock.removeItem).not.toHaveBeenCalled()
						}),
					},
				},
			} as any)

			await import('./api')
		})

		it('does not remove token on 401 for register endpoint', async () => {
			const error = {
				response: { status: 401 },
				config: { url: '/api/auth/register' },
			}

			mockedAxios.create.mockReturnValue({
				...mockedAxios,
				interceptors: {
					request: { use: vi.fn() },
					response: {
						use: vi.fn((_, errorHandler) => {
							errorHandler(error)
							expect(localStorageMock.removeItem).not.toHaveBeenCalled()
						}),
					},
				},
			} as any)

			await import('./api')
		})
	})
})