import { type JSX } from 'react'
import { Navigate } from 'react-router-dom'
import { useAuth } from '../contexts/AuthContext'

export function PrivateRoute({ children }: { children: JSX.Element }) {
	const { user, isLoading } = useAuth()

	if (isLoading) return <p>Carregando...</p>
	if (!user) return <Navigate to="/login" replace />
	return children
}

export function PublicRoute({ children }: { children: JSX.Element }) {
	const { user, isLoading } = useAuth()

	if (isLoading) return <p>Carregando...</p>
	if (user) return <Navigate to="/dashboard" replace />
	return children
}
