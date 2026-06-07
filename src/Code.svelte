<script>
  export let line = '';

  const copyToClipboard = (e) => {
    // Brittle, but correct:
    //   <img> =(sibling)=> <code> =(last child)=> <input>
    const cmdInput = e.target.previousElementSibling.lastChild;

    // The following was derived from
    // https://gist.github.com/rproenca/64781c6a1329b48a455b645d361a9aa3

    const iOS = navigator.userAgent.match(/ipad|iphone/i);

    if (iOS) {
      const range = document.createRange();
      range.selectNodeContents(cmdInput);
      selection = window.getSelection();
      selection.removeAllRanges();
      selection.addRange(range);
    } else {
      cmdInput.select();
    }

    cmdInput.setSelectionRange(0, 999999);

    const success = document.execCommand('copy');

    if (success) {
      console.log('Copied this command to clipboard!', cmdInput.value);
    } else {
      console.log('Copying command to clipboard failed :-(');
    }
  }
</script>

<style>
  .code-ctn {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
    color: #e6edf3;
    background-color: var(--bg-code);
    height: 38px;
    border-radius: 6px;
    padding: 0 16px;
    max-width: 100%;
    border: 1px solid #30363d;
    font-family: "IBM Plex Mono", monospace;
  }

  .code-ctn img {
    margin-left: 40px;
    cursor: pointer;
    opacity: 0.7;
    transition: opacity 0.15s ease;
  }

  .code-ctn img:hover {
    opacity: 1;
  }

  code input {
    color: #e6edf3;
    background-color: var(--bg-code);
    border: none;
    min-width: 200px;
    font-family: "IBM Plex Mono", monospace;
    font-style: normal;
    font-weight: normal;
    font-size: 14px;
    line-height: 18px;
  }
</style>

<div style="display: flex; min-width: 0;">
  <div class="code-ctn">
    <code>
      $&nbsp;
      <input type="text" value={line} />
    </code>
    <img
      src="/img/copy-to-clipboard.svg"
      alt="Copy to Clipboard Icon"
      on:click={copyToClipboard}
    >
  </div>
</div>
