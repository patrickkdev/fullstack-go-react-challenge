import { useEffect, useState } from 'react'
import { useAuth } from '../../contexts/AuthContext'
import { jobService } from '../../services/jobService'
import type { Job } from '../../types'

export default function MyJobsPage() {
	const { user } = useAuth()
	const [myJobs, setMyJobs] = useState<Job[]>([])
	const [loading, setLoading] = useState(false)
	const [error, setError] = useState('')

	const loadMyJobs = async () => {
		if (!user) return
		setLoading(true)
		setError('')
		try {
			const data = await jobService.getMyJobs()
			setMyJobs(data)
		} catch (err: unknown) {
			setError(err instanceof Error ? err.message : 'Falha ao carregar suas vagas.')
		} finally {
			setLoading(false)
		}
	}

	useEffect(() => { loadMyJobs() }, [user])

	return (
		<>
			{loading && <p>Carregando...</p>}
			{error && <p className="error">{error}</p>}
			<ul className="job-list">
				{myJobs.length === 0 ? (
					<li className="empty">Nenhuma vaga criada.</li>
				) : (
					myJobs.map((job) => (
						<li key={job.id} className="job-card">
							<h3>{job.title}</h3>
							<p>{job.description}</p>
							<p>Empresa: {job.company ?? '—'}</p>
						</li>
					))
				)}
			</ul>
		</>
	)
}
