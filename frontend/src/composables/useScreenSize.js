// composables/useScreenSize.js
import { ref, computed, onMounted, onUnmounted } from 'vue';

export function useScreenSize() {
    const windowWidth = ref(0);
    
    const isMobile = computed(() => windowWidth.value < 768);

    const updateScreenSize = () => {
        windowWidth.value = window.innerWidth;
    };

    onMounted(() => {
        updateScreenSize();
        window.addEventListener('resize', updateScreenSize);
    });

    onUnmounted(() => {
        window.removeEventListener('resize', updateScreenSize);
    });

    return {
        isMobile
    };
}