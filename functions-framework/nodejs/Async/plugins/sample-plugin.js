const {Plugin} = require('../node_modules/@openfunction/functions-framework/build/src/openfunction/plugin');
class Sample extends Plugin{
    constructor(){
        super('sample','v1')
    }
    async execPreHook(ctx,data){
        console.log('-----------test pre-----------')
    }
    async execPostHook(ctx,data){
        console.log('----------- test post-----------')
    }
}

exports.Sample = Sample