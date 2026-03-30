import type { AuthCredentials, AuthResponse, AuthUser, RegisterCredentials } from '../types'
import { api } from './api'

export const authService = {
  login: async (credentials: AuthCredentials) => {
    const { data } = await api.post<AuthResponse>('/login', credentials)
    return data
  },
	register: async (credentials: RegisterCredentials) => {
		const { data } = await api.post<AuthResponse>('/register', credentials)
		return data
	},
	profile: async () => {
		const { data } = await api.get<AuthUser>('/me')
		return data
	},
}
