export function useHelpers() {
    const formatDate = (dateString) => {
        if (!dateString) return ''
        const date = new Date(dateString)
        return date.toLocaleDateString('uk-UA', {
            year: 'numeric',
            month: 'long',
            day: 'numeric'
        })
    }

    const formatMileage = (mileage) => {
        return new Intl.NumberFormat('uk-UA').format(mileage) + ' км'
    }

    const formatPrice = (price) => {
        if (!price) return '$0'
        return new Intl.NumberFormat('en-US', {
            style: 'currency',
            currency: 'USD',
            minimumFractionDigits: 0
        }).format(price)
    }

    const formatEngineVolume = (volume) => {
        if (!volume) return '0.0'
        return (volume / 10).toFixed(1)
    }


    return {
        formatDate,
        formatPrice,
        formatMileage,
        formatEngineVolume
    }
}