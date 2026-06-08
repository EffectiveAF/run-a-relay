<script>
  export let line = '';

  const copyToClipboard = (e) => {
    const cmdInput = e.target.previousElementSibling.lastChild;
    const iOS = navigator.userAgent.match(/ipad|iphone/i);

    if (iOS) {
      const range = document.createRange();
      range.selectNodeContents(cmdInput);
      const selection = window.getSelection();
      selection.removeAllRanges();
      selection.addRange(range);
    } else {
      cmdInput.select();
    }

    cmdInput.setSelectionRange(0, 999999);
    document.execCommand('copy');
  };
</script>

<style>
  .code-wrap {
    display: flex;
    margin: 8px 0;
  }

  .code-ctn {
    display: flex;
    align-items: center;
    justify-content: space-between;
    background: var(--bg-code);
    border: 1px solid rgba(255,255,255,0.08);
    border-radius: 8px;
    padding: 0 14px;
    height: 40px;
    width: 100%;
    max-width: 560px;
    font-family: "IBM Plex Mono", "Fira Code", "Cascadia Code", monospace;
    font-size: 13px;
    color: #c9d1d9;
  }

  code {
    display: flex;
    align-items: center;
    gap: 6px;
    flex: 1;
    min-width: 0;
  }

  .prompt {
    color: #58a6ff;
    flex-shrink: 0;
    font-weight: 600;
  }

  code input {
    flex: 1;
    background: transparent;
    border: none;
    color: #c9d1d9;
    font-family: inherit;
    font-size: inherit;
    outline: none;
    min-width: 0;
    padding: 0;
  }

  .copy-btn {
    background: transparent;
    border: none;
    box-shadow: none;
    height: auto;
    padding: 4px 6px;
    cursor: pointer;
    border-radius: 4px;
    flex-shrink: 0;
    margin-left: 8px;
    opacity: 0.5;
    transition: opacity 0.12s;
  }

  .copy-btn:hover { opacity: 1; background: transparent; }

  .copy-btn img {
    display: block;
    width: 14px;
    height: 14px;
    filter: invert(1);
  }
</style>

<div class="code-wrap">
  <div class="code-ctn">
    <code>
      <span class="prompt">$</span>
      <input type="text" value={line} readonly />
    </code>
    <button class="copy-btn" on:click={copyToClipboard} title="Copy to clipboard">
      <img src="/img/copy-to-clipboard.svg" alt="Copy" />
    </button>
  </div>
</div>
