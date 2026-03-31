import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { useAuth } from '../../contexts/AuthContext'
import { jobService } from '../../services/jobService'

export default function CreateJobPage() {
  const { user } = useAuth()
  const navigate = useNavigate()

  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')
  const [company, setCompany] = useState('')
  const [location, setLocation] = useState('')
  const [salary, setSalary] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')
  const [success, setSuccess] = useState('')

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    setError('')
    setSuccess('')

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
      await jobService.createJob({ title, description, company: company || undefined, location: location || undefined, salary: salary || undefined })
      setSuccess('Vaga criada com sucesso!')
      setTitle('')
      setDescription('')
      setCompany('')
      setLocation('')
      setSalary('')

      setTimeout(() => {
        navigate('/dashboard/my-jobs')
      }, 1200)
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : 'Falha ao criar vaga.')
    } finally {
      setLoading(false)
    }
  }

  return (
    <main className="create-job-page">
      <header className="page-header">
        <h1>Criar nova vaga</h1>
        <div className="header-actions">
          <Link to="/dashboard/my-jobs" className="button">Voltar para minhas vagas</Link>
          <Link to="/dashboard/jobs" className="button button-secondary">Ver todas as vagas</Link>
        </div>
      </header>

      <section className="job-form">
        <form onSubmit={handleSubmit}>
          <label>
            Título *
            <input value={title} onChange={(e) => setTitle(e.target.value)} placeholder="Ex.: Desenvolvedor Frontend" />
          </label>

          <label>
            Descrição *
            <textarea value={description} onChange={(e) => setDescription(e.target.value)} rows={4} placeholder="Descreva a vaga com responsabilidades, benefícios e habilidades" />
          </label>

          <label>
            Empresa
            <input value={company} onChange={(e) => setCompany(e.target.value)} placeholder="Ex.: Acme Corp" />
          </label>

          <label>
            Localização
            <input value={location} onChange={(e) => setLocation(e.target.value)} placeholder="Ex.: Remoto" />
          </label>

          <label>
            Salário
            <input value={salary} onChange={(e) => setSalary(e.target.value)} placeholder="Ex.: 6.000" />
          </label>

          <button type="submit" disabled={loading}>{loading ? 'Enviando...' : 'Criar vaga'}</button>

          {success && <p className="success">{success}</p>}
          {error && <p className="error">{error}</p>}
        </form>
      </section>
    </main>
  )
}
