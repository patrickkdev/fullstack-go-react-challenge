/* eslint-disable react-refresh/only-export-components */

import type { ReactNode } from 'react'
import { createContext, useContext, useEffect, useMemo, useState } from 'react'
import { authService } from '../services/authService'
import type { AuthCredentials, AuthUser, RegisterCredentials } from '../types'

type AuthContextType = {
	user: AuthUser | null
	isLoading: boolean
	error: string | null
	login: (credentials: AuthCredentials) => Promise<void>
	register: (credentials: RegisterCredentials) => Promise<void>
	logout: () => void
}

const AuthContext = createContext<AuthContextType | undefined>(undefined)

export function AuthProvider({ children }: { children: ReactNode }) {
	const [user, setUser] = useState<AuthUser | null>(null)
	const [isLoading, setIsLoading] = useState(true)
	const [error, setError] = useState<string | null>(null)

	useEffect(() => {
		const token = window.localStorage.getItem('authToken')
		if (!token) {
			setIsLoading(false)
			setUser(null)
			return
		}

		async function loadSession() {
			try {
				const user = await authService.profile()
				setUser(user)
			} catch {
				window.localStorage.removeItem('authToken')
				setUser(null)
			} finally {
				setIsLoading(false)
			}
		}

		loadSession()
	}, [])

	const login = async (credentials: AuthCredentials) => {
		setIsLoading(true)
		setError(null)
		try {
			const { user } = await authService.login(credentials)
			window.localStorage.setItem('authToken', user.sessionToken)
			setUser(user)
		} catch (err: unknown) {
			const message = err instanceof Error ? err.message : 'Falha ao fazer login'
			setError(message)
			throw err
		} finally {
			setIsLoading(false)
		}
	}

	const register = async (credentials: RegisterCredentials) => {
		setIsLoading(true)
		setError(null)
		try {
			const { user } = await authService.register(credentials)
			window.localStorage.setItem('authToken', user.sessionToken)
			setUser(user)
		} catch (err: unknown) {
			const message = err instanceof Error ? err.message : 'Falha ao cadastrar'
			setError(message)
			throw err
		} finally {
			setIsLoading(false)
		}
	}

	const logout = () => {
		window.localStorage.removeItem('authToken')
		setUser(null)
	}

	const value = useMemo(
		() => ({ user, isLoading, error, login, register, logout }),
		[user, isLoading, error],
	)

	return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
}

export function useAuth() {
	const context = useContext(AuthContext)
	if (!context) {
		throw new Error('useAuth must be used within AuthProvider')
	}
	return context
}
