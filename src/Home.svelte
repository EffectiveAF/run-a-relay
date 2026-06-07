<script>
  import { onMount } from 'svelte';
  import { Link } from 'svelte-routing';
  import { currentStepIndex, currentPath } from './stores.js';

  onMount(() => {
    currentPath.set('/');
    currentStepIndex.set(0);
  });
</script>

<style>
  .hero {
    padding: 64px 48px 48px;
    max-width: 680px;
  }

  .eyebrow {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    font-size: 0.75rem;
    font-weight: 600;
    letter-spacing: 0.08em;
    text-transform: uppercase;
    color: var(--accent-light);
    margin-bottom: 20px;
    background: var(--bg-active);
    padding: 4px 10px;
    border-radius: 100px;
  }

  .dot {
    width: 6px;
    height: 6px;
    border-radius: 50%;
    background: var(--accent);
    display: inline-block;
    animation: pulse 2s ease-in-out infinite;
  }

  @keyframes pulse {
    0%, 100% { opacity: 1; transform: scale(1); }
    50% { opacity: 0.5; transform: scale(0.8); }
  }

  h1 {
    font-size: 2.75rem;
    font-weight: 700;
    line-height: 1.1;
    letter-spacing: -0.04em;
    margin-bottom: 16px;
  }

  .subtitle {
    font-size: 1.0625rem;
    line-height: 1.65;
    color: var(--text-2);
    margin-bottom: 32px;
    max-width: 520px;
  }

  .cta-row {
    display: flex;
    align-items: center;
    gap: 12px;
    flex-wrap: wrap;
  }

  .btn-primary {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    height: 40px;
    padding: 0 20px;
    background: var(--accent);
    color: #fff;
    font-size: 0.875rem;
    font-weight: 500;
    border-radius: 8px;
    text-decoration: none;
    transition: filter 0.12s, box-shadow 0.12s;
    box-shadow: 0 4px 16px var(--accent-glow);
    white-space: nowrap;
  }

  .btn-primary:hover {
    color: #fff;
    filter: brightness(1.1);
    box-shadow: 0 6px 22px var(--accent-glow);
  }

  .btn-secondary {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    height: 40px;
    padding: 0 20px;
    background: var(--bg-raised);
    color: var(--text);
    font-size: 0.875rem;
    font-weight: 500;
    border-radius: 8px;
    border: 1px solid var(--border-strong);
    text-decoration: none;
    transition: background 0.12s;
  }

  .btn-secondary:hover {
    color: var(--text);
    background: var(--bg-subtle);
  }

  .divider {
    border: none;
    border-top: 1px solid var(--border);
    margin: 48px 0 0;
  }

  .section {
    padding: 0 48px 80px;
    max-width: 680px;
  }

  .section h2 {
    font-size: 1.125rem;
    font-weight: 600;
    letter-spacing: -0.01em;
    margin: 40px 0 12px;
  }

  .section h2:first-child { margin-top: 0; }

  .feature-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
    gap: 12px;
    margin-top: 32px;
  }

  .feature-card {
    background: var(--bg-raised);
    border: 1px solid var(--border);
    border-radius: 10px;
    padding: 20px;
    transition: border-color 0.15s;
  }

  .feature-card:hover { border-color: var(--border-strong); }

  .feature-icon {
    font-family: "IBM Plex Mono", monospace;
    font-size: 1rem;
    font-weight: 700;
    color: var(--accent-light);
    margin-bottom: 10px;
  }

  .feature-title {
    font-size: 0.875rem;
    font-weight: 600;
    color: var(--text);
    margin-bottom: 6px;
  }

  .feature-desc {
    font-size: 0.8125rem;
    color: var(--text-2);
    line-height: 1.55;
    margin: 0;
  }

  @media (max-width: 900px) {
    .hero, .section { padding-left: 24px; padding-right: 24px; }
    h1 { font-size: 2.125rem; }
  }

  @media (max-width: 640px) {
    .hero { padding-top: 40px; padding-left: 16px; padding-right: 16px; }
    .section { padding-left: 16px; padding-right: 16px; }
    h1 { font-size: 1.75rem; }
  }
</style>

<div class="hero">
  <div class="eyebrow">
    <span class="dot"></span>
    Free & Open Source
  </div>

  <h1>Run a Tor relay.<br />Help keep the internet free.</h1>

  <p class="subtitle">
    A step-by-step guide to setting up your own Tor relay and contributing
    to a more private, open internet — in under 30 minutes.
  </p>

  <div class="cta-row">
    <a class="btn-primary" href="/step/intro">Get started →</a>
    <a
      class="btn-secondary"
      href="https://github.com/EffectiveAF/run-a-relay"
      target="_blank"
      rel="noopener noreferrer"
    >GitHub ↗</a>
  </div>
</div>

<hr class="divider" />

<div class="section">
  <h2>What is this?</h2>
  <p>
    Run A Relay makes it fast and easy to set up a new
    <a href="https://www.torproject.org/" target="_blank" rel="noopener noreferrer">Tor</a>
    relay, with special emphasis on helping you choose a server provider based on
    <a href="https://trac.torproject.org/projects/tor/wiki/doc/GoodBadISPs" target="_blank" rel="noopener noreferrer">
      crowdsourced information from The Tor Project
    </a>.
    We've distilled everything down into actionable steps you can follow in one sitting.
  </p>

  <div class="feature-grid">
    <div class="feature-card">
      <div class="feature-icon">→</div>
      <div class="feature-title">Fast setup</div>
      <p class="feature-desc">From zero to running relay in under 30 minutes.</p>
    </div>
    <div class="feature-card">
      <div class="feature-icon">#</div>
      <div class="feature-title">Curated hosts</div>
      <p class="feature-desc">Tor-friendly server providers, community-vetted.</p>
    </div>
    <div class="feature-card">
      <div class="feature-icon">*</div>
      <div class="feature-title">Real impact</div>
      <p class="feature-desc">Each relay helps thousands protect their privacy daily.</p>
    </div>
  </div>

  <h2>Before you start</h2>
  <p>
    You'll need a server running Linux (Debian or Ubuntu preferred), at least
    $10/month to spare for hosting, and a basic comfort level with the terminal.
    If you'd prefer to donate instead, the guide covers that option too.
  </p>
</div>
