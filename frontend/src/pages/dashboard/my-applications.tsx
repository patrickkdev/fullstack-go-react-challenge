import { useCallback, useEffect, useState } from 'react'
import { useAuth } from '../../contexts/AuthContext'
import { jobService } from '../../services/jobService'
import type { Application } from '../../types'

export default function MyApplicationsPage() {
	const { user } = useAuth()
	const [applications, setApplications] = useState<Application[]>([])
	const [loading, setLoading] = useState(false)
	const [error, setError] = useState('')

	const loadApplications = useCallback(async () => {
		if (!user) return
		setLoading(true)
		setError('')
		try {
			const data = await jobService.getApplications()
			setApplications(data)
		} catch (err: unknown) {
			setError(err instanceof Error ? err.message : 'Falha ao carregar candidaturas.')
		} finally {
			setLoading(false)
		}
	}, [user])

	useEffect(() => { loadApplications() }, [loadApplications])

	return (
		<>
			{loading && <p>Carregando...</p>}
			{error && <p className="error">{error}</p>}
			<ul className="job-list">
				{applications.length === 0 ? (
					<li className="empty">Nenhuma candidatura registrada.</li>
				) : (
					applications.map((app) => (
						<li key={app.id} className="job-card">
							<h3>{`Vaga #${app.jobId}`}</h3>
							<p>Status: {app.status}</p>
							<p>Criado em: {new Date(app.createdAt).toLocaleString()}</p>
						</li>
					))
				)}
			</ul>
		</>
	)
}
