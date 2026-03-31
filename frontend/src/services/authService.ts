import type { AuthCredentials, AuthResponse, AuthUser, RegisterCredentials } from '../types'
import { api } from './api'

export const authService = {
	login: async (credentials: AuthCredentials) => {
    try {
      const { data } = await api.post<AuthResponse>('/login', credentials)
      return data
    } catch (err: unknown) {
      const message =
      err instanceof Error
        ? err.message
        : (typeof err === 'object' && err !== null && 'response' in err
          ? (err as { response?: { data?: { message?: string } } }).response?.data?.message
          : undefined) || 'Email ou senha inválidos'

      throw new Error(message)
    }
  },
	register: async (credentials: RegisterCredentials) => {
		const { data } = await api.post<AuthResponse>('/register', credentials)
		return data
	},
	profile: async () => {
		const { data } = await api.get<{ user: AuthUser }>('/me')
		return data.user
	},
}
