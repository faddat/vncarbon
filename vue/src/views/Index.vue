<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form path="faddat.vncarbon.vncarbon" type="vncarbon" :fields="[ ['creator', 1, 'string'] , ['carbon', 2, 'int32'] , ['emitter', 3, 'string'] , ['date', 4, 'string'] ]" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
