import { useCallback, useEffect, useState } from 'react'
import { useAuth } from '../../contexts/AuthContext'
import { jobService } from '../../services/jobService'
import type { Job } from '../../types'

export default function JobsPage() {
	const { user } = useAuth()
	const [jobs, setJobs] = useState<Job[]>([])
	const [search, setSearch] = useState('')
	const [loading, setLoading] = useState(false)
	const [error, setError] = useState('')

	const loadJobs = useCallback(async () => {
		setLoading(true)
		setError('')
		try {
			const data = await jobService.getJobs(search)
			setJobs(data)
		} catch (err: unknown) {
			setError(err instanceof Error ? err.message : 'Falha ao carregar vagas.')
		} finally {
			setLoading(false)
		}
	}, [search])

	useEffect(() => {
		loadJobs()
	}, [loadJobs])

	const handleApply = async (jobId: number) => {
		if (!user) {
			alert('Usuário não autenticado.')
			return
		}

		try {
			await jobService.applyJob(jobId)
			alert('Candidatado com sucesso!')
			loadJobs()
		} catch (err: unknown) {
			alert(err instanceof Error ? err.message : 'Falha na candidatura.')
		}
	}

	return (
		<>
			<section className="search-row">
				<input value={search} onChange={(e) => setSearch(e.target.value)} placeholder="Buscar vagas" />
				<button onClick={loadJobs}>Atualizar</button>
			</section>

			{loading && <p>Carregando...</p>}
			{error && <p className="error">{error}</p>}

			<ul className="job-list">
				{jobs.length === 0 ? (
					<li className="empty">Nenhuma vaga encontrada.</li>
				) : (
					jobs.map((job) => (
						<li className="job-card" key={job.id}>
              <p className="job-id">ID: {job.id}</p>
							<h3>{job.title}</h3>
							<p>{job.description}</p>
							<p>Empresa: {job.company ?? '—'}</p>
							<button onClick={() => handleApply(job.id)}>Candidatar-se</button>
						</li>
					))
				)}
			</ul>
		</>
	)
}
