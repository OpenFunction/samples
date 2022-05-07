export const tryKnative = (req, res) => {
  res.send(`Hello, ${req.query.u || 'World'}!`);
};

export const tryAsync = (ctx, data) => {
  console.log('Data received: %o', data);
  ctx.send(data);
};
