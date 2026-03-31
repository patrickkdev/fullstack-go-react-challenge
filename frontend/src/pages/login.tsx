import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAuth } from '../contexts/AuthContext'

export default function Login() {
	const [email, setEmail] = useState('')
	const [password, setPassword] = useState('')
	const [error, setError] = useState('')
	const [loading, setLoading] = useState(false)
	const { login } = useAuth()
	const navigate = useNavigate()

	const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
		event.preventDefault()
		setError('')

		if (!email || !password) {
			setError('Email e senha são obrigatórios.')
			return
		}

		if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
			setError('Email inválido.')
			return
		}

		setLoading(true)
		try {
			await login({ email, password })
			navigate('/dashboard')
		} catch (err: unknown) {
      const message = err instanceof Error ? err.message : 'Falha no login.'
			setError(message)
		} finally {
      setLoading(false)
		}
	}

	return (
		<main className="auth-page">
			<h1>Login</h1>
			<form onSubmit={handleSubmit} className="auth-form">
				<label>
					Email
					<input value={email} onChange={(e) => setEmail(e.target.value)} type="email" autoComplete="email" />
				</label>
				<label>
					Senha
					<input value={password} onChange={(e) => setPassword(e.target.value)} type="password" autoComplete="current-password" />
				</label>
				<button disabled={loading} type="submit">
					{loading ? 'Entrando...' : 'Entrar'}
				</button>
				{error && <p className="error">{error}</p>}
			</form>
      <p>Ainda não possui uma conta? <a href="/register">Registrar</a></p>
		</main>
	)
}