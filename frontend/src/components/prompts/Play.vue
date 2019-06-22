<template>
    <div class="card floating" id="share">
        <div class="card-title">
            <h2>{{ $t('prompts.title') }}</h2>
        </div>

        <div class="card-content">
            <p>{{ $t('prompts.playPrompt') }}</p>
        </div>

        <div class="card-action">
            <button class="button button--flat"
                    @click="$store.commit('closeHovers')"
                    :aria-label="$t('buttons.close')"
                    :title="$t('buttons.close')">{{ $t('buttons.close') }}</button>
        </div>
    </div>
</template>

<script>
export default { name: 'play' }
</script>