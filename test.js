const parse = require('./');
const json = JSON.stringify(parse('{{ . }}'));
if (json !== '{"type":"ListNode","pos":0,"nodes":[{"type":"ActionNode","pos":3,"pipe":{"type":"PipeNode","pos":3,"cmds":[{"type":"CommandNode","pos":3,"args":[{"type":"DotNode","pos":3}]}],"decl":[]}}]}') {
  console.info('Fail');
  process.exit(1);
} else {
  console.info('Success');
}
