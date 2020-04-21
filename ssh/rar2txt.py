import re

rar_orig = open('../src/Steps.svelte').read()

on_step = 1

print('# RunARelay.org Tutorial')

for title, p in re.findall(r'(?:title="(.*?)")|(?:<p>(.*?)</p>)', rar_orig, re.DOTALL):
    if title:
      if on_step > 1:
          print('---')
      print('\n## Step ' + str(on_step) + ':', title + '\n')
      on_step += 1

    if p and '<button' not in p and '{@html' not in p:
      p = re.sub(r'\n +', '\n ', p.lstrip())
      p = re.sub(r'<ExternalLink href="(.*?)">\s*(.*?)\s*</ExternalLink>', r'\2 ( \1 )', p)
      print(p.
            replace('&nbsp;', ' ').
            replace('  ', ' ').
            replace('<em>', '**').
            replace('</em>', '**')
      )
