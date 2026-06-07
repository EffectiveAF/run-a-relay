<script>
  import { Link } from 'svelte-routing';
  import ExternalLink from './ExternalLink.svelte';
</script>

<style>

  footer {
    display: flex;
    flex-direction: row;
    position: absolute;
    right: 0;
    bottom: 0;
    left: 0;
    height: 24px;
    padding: 16px;
    margin: 0 0 12px 0;
    justify-content: space-between;
    margin-top: 60vh;
    border-top: 1px solid var(--border-color);
    color: var(--text-secondary);
  }

  span {
    content: "\2764";
  }

  footer ul, footer > .footer-middle {
    display: flex;
    flex: 1;
  }

  footer ul, footer ul li {
    display: flex;
    flex-direction: row;
    font-size: 12px;
    align-items: center;
    color: var(--text-secondary);
  }

  .icon-github {
    display: inline-block;
    width: 25px;
    height: 24px;
    background-color: var(--text-secondary);
    mask: url(/img/github_logo.svg) no-repeat center / contain;
    -webkit-mask: url(/img/github_logo.svg) no-repeat center / contain;
    transition: background-color 0.15s ease;
  }

  .icon-github:hover {
    background-color: var(--text-primary);
  }

  .icon-tor {
    display: inline-block;
    width: 25px;
    height: 24px;
    background-color: var(--text-secondary);
    mask: url(/img/tor.svg) no-repeat center / contain;
    -webkit-mask: url(/img/tor.svg) no-repeat center / contain;
    transition: background-color 0.15s ease;
  }

  .icon-tor:hover {
    background-color: var(--accent-primary);
  }

  footer ul li:not(:first-child) {
    padding-left: 24px;
  }

  footer > .footer-middle {
    justify-content: center;
  }

  .list-end {
    justify-content: flex-end;
  }

  @media (max-width: 1200px) {
    footer {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      margin-top: 50vh;
    }

    footer ul {
      padding-top: 24px;
    }

    footer ul:last-child {
      padding-bottom: 24px;
    }

    footer a img {
      margin: 0;
    }

    footer > .footer-middle {
      order: -1;
    }

    .list-end {
      justify-content: center;
    }

    @media (max-width: 600px) {
      .hide-xsmall {
        display: none;
      }
    }
  }

</style>

<footer>

  <ul>
    <li>© 2020&nbsp;<ExternalLink href="https://effective.af/">EffectiveAF</ExternalLink></li>

    <li><Link to="/terms">Terms</Link></li>

    <li><Link to="/privacy"><li>Privacy</li></Link></li>

    <li class="hide-xsmall"><Link to="/about">About</Link></li>

    <li class="hide-xsmall"><ExternalLink href="https://donate.torproject.org/">Donate to Tor</ExternalLink></li>
  </ul>

  <div class="footer-middle">
    <ExternalLink href="https://github.com/EffectiveAF/run-a-relay"><div class="icon-github" role="img" aria-label="GitHub"></div></ExternalLink>
    <div style="width: 8px;"></div>
    <ExternalLink href="http://vxw5mjyap4bgb4w7h7qkeknxszmgp5efjtg77mnwjtbczynihsij3myd.onion"><div class="icon-tor" role="img" aria-label="Visit on Tor"></div></ExternalLink>
  </div>

  <ul class="list-end">
    <li>Made by&nbsp;<ExternalLink href="https://github.com/elimisteve">Steve Phillips</ExternalLink>&nbsp;&&nbsp;<ExternalLink href="https://github.com/timsully">Tim Sullivan</ExternalLink>&nbsp;with&nbsp;<span>&#10084;</span>
    </li>
  </ul>

</footer>
