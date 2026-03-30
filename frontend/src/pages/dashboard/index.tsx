import { Link, Outlet } from 'react-router-dom'
import { useAuth } from '../../contexts/AuthContext'

export default function DashboardLayout() {
	const { user, logout } = useAuth()

	return (
		<main className="dashboard-page">
			<header className="dashboard-header">
				<div>
					<h1>Bem-vindo, {user?.email}</h1>
					<p>Área interna de vagas</p>
				</div>
				<button onClick={logout}>Sair</button>
			</header>

			<nav className="tabs">
				<Link to="jobs" className="tab-link">
					Vagas gerais
				</Link>
				<Link to="my-jobs" className="tab-link">
					Minhas vagas
				</Link>
				<Link to="my-applications" className="tab-link">
					Minhas candidaturas
				</Link>
			</nav>

			<Outlet />
		</main>
	)
}
