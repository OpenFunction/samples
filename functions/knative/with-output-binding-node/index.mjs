export const tryKnative = (req, res) => {
  res.send(`Hello, ${req.query.u || 'World'}!`);
};

export const tryKnativeAsync = async (ctx, data) => {
  console.log('Data received: %o', data);
  await ctx.send(data);

  // Optional to send ANY data back as HTTP response
  // Request data is also accessible via "ctx.req"
  // ctx.res.send(ctx.req.query);
};
