<script>
  import { Link } from 'svelte-routing';
  import ExternalLink from './ExternalLink.svelte';
</script>

<style>

  footer {
    display: flex;
    flex-direction: row;
    height: 24px;
    margin: 24px 16px 0px 16px;
    justify-content: space-between;
    margin-top: 60vh;
  }

  footer > ul {
    min-width: 40vw;
  }

  span {
    content: "\2764";
  }

  footer ul {
    flex-grow: 1;
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

  .list-end {
    justify-content: flex-end;
  }

  @media (max-width: 960px) {
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

    footer > :global(a) {
      order: -1;
    }

    .list-end {
      justify-content: center;
    }

  }

</style>

<footer>

  <ul>
    <li>© 2020 Run A Relay</li>

    <li><Link to="/terms">Terms</Link></li>

    <li><Link to="/privacy"><li>Privacy</li></Link></li>

    <li><ExternalLink href="https://donate.torproject.org/">Donate to Tor</ExternalLink></li>

    <li><Link to="/about">About</Link></li>
  </ul>

  <ExternalLink href="https://github.com/elimisteve/run-a-relay"><img src="./img/github_logo.svg" alt="Github Logo" width="25px" height="24px"></ExternalLink>

  <ul class="list-end">
    <li>Made by&nbsp;<ExternalLink href="https://github.com/elimisteve">Steve Phillips</ExternalLink>&nbsp;&&nbsp;<ExternalLink href="https://github.com/timsully">Tim Sullivan&nbsp;</ExternalLink>&nbsp;with&nbsp;<span>&#10084;</span>
    </li>
  </ul>

</footer>