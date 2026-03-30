export default function LandingPage() {
	return (
		<main className="landing-page">
			<section className="hero">
				<div className="hero-content">
					<h1>Recrutamento e Seleção</h1>
					<p>Backend em Go e Frontend em React + Vite.</p>
					<div style={{ display: 'flex', alignItems: 'center', gap: '0.6rem' }}>
						<a className="cta-button" href="/login">Entrar</a>
						<a className="cta-button" href="/register">Registrar</a>
					</div>
				</div>
			</section>

			<section id="features" className="features">
				<h2>Features</h2>
				<div className="feature-grid">
					<article className="feature-card">
						<h3>Tela de Cadastro (Email + Senha)</h3>
					</article>
					<article className="feature-card">
            <h3>Tela de Login (Email + Senha)</h3>
					</article>
					<article className="feature-card">
						<h3>Tela quando logado (Minhas vagas / Minhas candidaturas)</h3>
					</article>
          <article className="feature-card">
            <h3>Registro de Vagas</h3>
          </article>
          <article className="feature-card">
            <h3>Cadastro em Vaga</h3>
          </article>
          <article className="feature-card">
            <h3>Busca de Vagas</h3>
          </article>
        </div>
      </section>
				
			<footer className="landing-footer">
				<p>© {new Date().getFullYear()} Recrutamento e Seleção</p>
			</footer>
		</main>
	)
}
