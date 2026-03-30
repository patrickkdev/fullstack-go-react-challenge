import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAuth } from '../contexts/AuthContext'

export default function Register() {
  const [name, setName] = useState('')
	const [email, setEmail] = useState('')
	const [password, setPassword] = useState('')
	const [error, setError] = useState('')
	const [loading, setLoading] = useState(false)

	const { register } = useAuth()
	const navigate = useNavigate()

	const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
		event.preventDefault()
		setError('')

		if (!name || !email || !password) {
			setError('Nome, email e senha são obrigatórios.')
			return
		}

    if (!/^[A-Za-z\s]+$/.test(name)) {
      setError('Nome deve conter apenas letras e espaços.')
      return
    }

		if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
			setError('Email inválido.')
			return
		}

		if (password.length < 6) {
			setError('Senha deve ter ao menos 6 caracteres.')
			return
		}

		setLoading(true)
		try {
			await register({ name, email, password })
			navigate('/dashboard')
		} catch (err: unknown) {
			const message = err instanceof Error ? err.message : 'Falha no cadastro.'
			setError(message)
		} finally {
			setLoading(false)
		}
	}

	return (
		<main className="auth-page">
			<h1>Registro</h1>
			<form onSubmit={handleSubmit} className="auth-form">
        <label>
					Nome
					<input value={name} onChange={(e) => setName(e.target.value)} type="text" autoComplete="name" />
				</label>
				<label>
					Email
					<input value={email} onChange={(e) => setEmail(e.target.value)} type="email" autoComplete="email" />
				</label>
				<label>
					Senha
					<input value={password} onChange={(e) => setPassword(e.target.value)} type="password" autoComplete="new-password" />
				</label>
				<button disabled={loading} type="submit">
					{loading ? 'Cadastrando...' : 'Cadastrar'}
				</button>
				{error && <p className="error">{error}</p>}
			</form>
      <p>Já possui uma conta? <a href="/login">Entrar</a></p>
		</main>
	)
}