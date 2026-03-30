import { useEffect, useState } from 'react'
import { jobService } from '../../services/jobService'
import type { Application } from '../../types'

export default function MyApplicationsPage() {
	const [applications, setApplications] = useState<Application[]>([])
	const [loading, setLoading] = useState(false)
	const [error, setError] = useState('')

	const loadApplications = async () => {
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
	}

	useEffect(() => { loadApplications() }, [])

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
							<h3>{app.job?.title ?? `Vaga #${app.jobId}`}</h3>
							<p>Status: {app.status}</p>
							<p>Criado em: {new Date(app.createdAt).toLocaleString()}</p>
						</li>
					))
				)}
			</ul>
		</>
	)
}
