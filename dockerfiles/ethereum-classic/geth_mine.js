var debounce, eth, now;

now = Date.now || function() {
  return (new Date).getTime();
};

debounce = function(func, wait) {
  var args, context, later, result, timeout, timestamp;
  timeout = void 0;
  args = void 0;
  context = void 0;
  timestamp = void 0;
  result = void 0;
  later = function() {
    var last;
    last = now() - timestamp;
    if (last < wait && last >= 0) {
      timeout = setTimeout(later, wait - last);
    } else {
      timeout = null;
      result = func.apply(context, args);
      if (!timeout) {
        context = args = null;
      }
    }
  };
  return function() {
    context = this;
    args = arguments;
    timestamp = now();
    if (!timeout) {
      timeout = setTimeout(later, wait);
    }
    return result;
  };
};

eth = web3.eth;

(function() {
  var config, ifNoPendingTransactions, ifNoPendingTransactionsDeb, ifNotMiningDo, ifNotMiningDoDeb, main, miner_start, miner_stop, old_geth, pendingTransactions, startTransactionMining, stillMining;
  console.log('geth_mine.js: start');
  console.log("node infos: " + (JSON.stringify(admin.nodeInfo)));
  config = {
    threads: 8
  };
  old_geth = admin.nodeInfo.name.match(/^Geth\/v1\.3/);
  miner_stop = function() {
    if (old_geth) {
      return miner.stop(config.threads);
    } else {
      return miner.stop();
    }
  };
  miner_start = function() {
    if (old_geth) {
      return miner.start(config.threads);
    } else {
      return miner.start();
    }
  };
  main = function() {
    miner_stop();
    return startTransactionMining();
  };
  pendingTransactions = function() {
    if (!eth.pendingTransactions) {
      return txpool.status.pending || txpool.status.queued;
    } else if (typeof eth.pendingTransactions === 'function') {
      return eth.pendingTransactions().length > 0;
    } else {
      return eth.pendingTransactions.length > 0 || eth.getBlock('pending').transactions.length > 0;
    }
  };
  ifNoPendingTransactions = function(callback) {
    if (!pendingTransactions()) {
      return callback();
    }
  };
  stillMining = function(callback) {
    return miner.hashrate > 0;
  };
  ifNotMiningDo = function(callbackWhenInactive) {
    if (!stillMining()) {
      return callbackWhenInactive();
    }
  };
  ifNoPendingTransactionsDeb = debounce(ifNoPendingTransactions, 200);
  ifNotMiningDoDeb = debounce(ifNotMiningDo, 200);
  startTransactionMining = function() {
    eth.filter('pending').watch(function() {
      return ifNotMiningDoDeb(function() {
        console.log('== Pending transactions! Looking for next block...');
        return miner.start(config.threads);
      });
    });
    return eth.filter('latest').watch(function() {
      return ifNoPendingTransactionsDeb(function() {
        console.log('== No transactions left. Stopping miner...');
        return miner.stop();
      });
    });
  };
  return main();
})();

// ---
// Original file https://raw.githubusercontent.com/makevoid/ethereum-geth-dev/master/lib/geth_mine.coffee
// generated by coffee-script 1.9.2