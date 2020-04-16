<script>
  import { Link } from 'svelte-routing';
  import ExternalLink from './ExternalLink.svelte';
</script>

<style>

  footer {
    display: flex;
    flex-direction: row;
    height: 24px;
    margin: 24px 16px 12px 16px;
    justify-content: space-between;
    margin-top: 60vh;
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
    <ExternalLink href="https://github.com/EffectiveAF/run-a-relay"><img src="/img/github_logo.svg" alt="Github Logo" width="25px" height="24px"></ExternalLink>
  </div>

  <ul class="list-end">
    <li>Made by&nbsp;<ExternalLink href="https://github.com/elimisteve">Steve Phillips</ExternalLink>&nbsp;&&nbsp;<ExternalLink href="https://github.com/timsully">Tim Sullivan</ExternalLink>&nbsp;with&nbsp;<span>&#10084;</span>
    </li>
  </ul>

</footer>
