export function useFormatters() {
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

    return {
        formatDate,
        formatPrice,
        formatNumber,
        formatMileage
    }
}