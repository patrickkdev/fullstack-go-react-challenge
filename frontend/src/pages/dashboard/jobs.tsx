import { useCallback, useEffect, useMemo, useState } from 'react'
import { useAuth } from '../../contexts/AuthContext'
import { jobService } from '../../services/jobService'
import type { Job } from '../../types'

function JobForm({ onCreated }: { onCreated: () => void }) {
	const [title, setTitle] = useState('')
	const [description, setDescription] = useState('')
	const [loading, setLoading] = useState(false)
	const [error, setError] = useState('')

	const { user } = useAuth()

	const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
		event.preventDefault()
		setError('')

		if (!title.trim() || !description.trim()) {
			setError('Título e descrição são obrigatórios.')
			return
		}

		if (!user) {
			setError('Usuário não autenticado.')
			return
		}

		setLoading(true)
		try {
			await jobService.createJob({ title, description })
			setTitle('')
			setDescription('')
			onCreated()
		} catch (err: unknown) {
			setError(err instanceof Error ? err.message : 'Falha ao criar vaga.')
		} finally {
			setLoading(false)
		}
	}

	return (
		<section className="job-form">
			<h2>Cadastrar vaga</h2>
			<form onSubmit={handleSubmit}>
				<label>
					Título
					<input value={title} onChange={(e) => setTitle(e.target.value)} />
				</label>
				<label>
					Descrição
					<textarea value={description} onChange={(e) => setDescription(e.target.value)} rows={3} />
				</label>
				<button type="submit" disabled={loading}>{loading ? 'Enviando...' : 'Cadastrar'}</button>
				{error && <p className="error">{error}</p>}
			</form>
		</section>
	)
}

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

	const filteredJobs = useMemo(() => {
		if (!search.trim()) return jobs
		return jobs.filter((job) => job.title.toLowerCase().includes(search.toLowerCase()))
	}, [jobs, search])

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
			<JobForm onCreated={loadJobs} />

			{loading && <p>Carregando...</p>}
			{error && <p className="error">{error}</p>}

			<ul className="job-list">
				{filteredJobs.length === 0 ? (
					<li className="empty">Nenhuma vaga encontrada.</li>
				) : (
					filteredJobs.map((job) => (
						<li className="job-card" key={job.id}>
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
