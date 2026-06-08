<script>
  import { onMount } from 'svelte';
  import { navigate } from 'svelte-routing';
  import { currentStepIndex, currentPath } from './stores.js';

  onMount(() => {
    currentPath.set('/');
    currentStepIndex.set(0);
  });
</script>

<style>
  .page-content { display: flex; flex-direction: column; gap: 0; }

  /* ── Hero ── */
  .hero { margin-bottom: 40px; }

  h1 { margin-bottom: 14px; }

  .subtitle {
    font-size: 0.9375rem;
    line-height: 1.6;
    color: var(--text-2);
    margin: 0 0 28px;
  }

  .cta-row {
    display: flex;
    align-items: center;
    gap: 10px;
    flex-wrap: wrap;
  }

  .btn-primary {
    display: inline-flex;
    align-items: center;
    gap: 5px;
    height: 36px;
    padding: 0 16px;
    background: var(--accent);
    color: #fff;
    font-size: 0.8125rem;
    font-weight: 500;
    border-radius: 7px;
    text-decoration: none;
    transition: filter 0.1s, box-shadow 0.1s;
    box-shadow: 0 2px 10px var(--accent-glow);
  }

  .btn-primary:hover {
    color: #fff;
    filter: brightness(1.1);
    box-shadow: 0 4px 16px var(--accent-glow);
  }

  .btn-secondary {
    display: inline-flex;
    align-items: center;
    gap: 5px;
    height: 36px;
    padding: 0 16px;
    background: var(--bg-raised);
    color: var(--text);
    font-size: 0.8125rem;
    font-weight: 500;
    border-radius: 7px;
    border: 1px solid var(--border-strong);
    text-decoration: none;
    transition: background 0.1s;
  }

  .btn-secondary:hover { color: var(--text); background: var(--bg-subtle); }

  /* ── Divider ── */
  .section-divider {
    border: none;
    border-top: 1px solid var(--border);
    margin: 0 0 36px;
  }

  /* ── Body section ── */
  .body-section { }

  /* ── Feature cards ── */
  .feature-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 10px;
    margin: 20px 0 32px;
  }

  @media (max-width: 700px) {
    .feature-grid { grid-template-columns: 1fr; }
  }

  .feature-card {
    background: var(--bg-raised);
    border: 1px solid var(--border);
    border-radius: 9px;
    padding: 18px;
    transition: border-color 0.12s;
  }

  .feature-card:hover { border-color: var(--border-strong); }

  .feature-icon {
    font-family: "IBM Plex Mono", monospace;
    font-size: 0.75rem;
    font-weight: 700;
    color: var(--accent-light);
    margin-bottom: 10px;
    display: block;
  }

  .feature-title {
    font-size: 0.8125rem;
    font-weight: 600;
    color: var(--text);
    margin-bottom: 5px;
  }

  .feature-desc {
    font-size: 0.8rem;
    color: var(--text-2);
    line-height: 1.5;
    margin: 0;
  }
</style>

<div class="page-content">

  <div class="hero">
    <h1>Run a Tor relay.<br />Help keep the internet free.</h1>
    <p class="subtitle">
      A step-by-step guide to setting up your own Tor relay and
      contributing to a more private, open internet — in under 30 minutes.
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

  <hr class="section-divider" />

  <div class="body-section">
    <h2>What is this?</h2>
    <p>
      Run A Relay makes it fast and easy to set up a new
      <a href="https://www.torproject.org/" target="_blank" rel="noopener noreferrer">Tor</a>
      relay, with special emphasis on helping you choose a server provider based on
      <a href="https://trac.torproject.org/projects/tor/wiki/doc/GoodBadISPs" target="_blank" rel="noopener noreferrer">
        crowdsourced information from The Tor Project
      </a>.
      Everything is distilled into clear steps you can follow in one sitting.
    </p>

    <div class="feature-grid">
      <div class="feature-card">
        <span class="feature-icon">→ fast</span>
        <div class="feature-title">Quick setup</div>
        <p class="feature-desc">From zero to running relay in under 30 minutes.</p>
      </div>
      <div class="feature-card">
        <span class="feature-icon"># hosts</span>
        <div class="feature-title">Curated hosts</div>
        <p class="feature-desc">Tor-friendly providers, community-vetted.</p>
      </div>
      <div class="feature-card">
        <span class="feature-icon">* impact</span>
        <div class="feature-title">Real impact</div>
        <p class="feature-desc">Each relay helps thousands protect their privacy daily.</p>
      </div>
    </div>

    <h2>Before you start</h2>
    <p>
      You'll need a Linux server (Debian or Ubuntu preferred), at least $10/month
      for hosting, and basic comfort with the terminal. If you'd prefer to donate
      instead of running a relay, the guide covers that option too.
    </p>
  </div>

</div>
