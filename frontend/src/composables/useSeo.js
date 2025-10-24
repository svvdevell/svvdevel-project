// src/composables/useSeo.js
// Композабл для оновлення meta tags в runtime (для користувачів)

export function useSeo() {
    const setMeta = (metaData) => {
        // Title
        if (metaData.title) {
            document.title = metaData.title;
        }

        // Description
        updateMetaTag('name', 'description', metaData.description);

        // Keywords
        if (metaData.keywords) {
            updateMetaTag('name', 'keywords', metaData.keywords);
        }

        // Open Graph
        if (metaData.title) {
            updateMetaTag('property', 'og:title', metaData.title);
        }
        if (metaData.description) {
            updateMetaTag('property', 'og:description', metaData.description);
        }
        if (metaData.ogImage) {
            updateMetaTag('property', 'og:image', metaData.ogImage);
        }
        if (metaData.url) {
            updateMetaTag('property', 'og:url', metaData.url);
        }

        // Twitter
        if (metaData.title) {
            updateMetaTag('name', 'twitter:title', metaData.title);
        }
        if (metaData.description) {
            updateMetaTag('name', 'twitter:description', metaData.description);
        }
        if (metaData.ogImage) {
            updateMetaTag('name', 'twitter:image', metaData.ogImage);
        }

        // Canonical
        if (metaData.url) {
            updateCanonical(metaData.url);
        }
    };

    const updateMetaTag = (attribute, key, content) => {
        if (!content) return;

        let element = document.querySelector(`meta[${attribute}="${key}"]`);

        if (element) {
            element.setAttribute('content', content);
        } else {
            element = document.createElement('meta');
            element.setAttribute(attribute, key);
            element.setAttribute('content', content);
            document.head.appendChild(element);
        }
    };

    const updateCanonical = (url) => {
        let element = document.querySelector('link[rel="canonical"]');

        if (element) {
            element.setAttribute('href', url);
        } else {
            element = document.createElement('link');
            element.setAttribute('rel', 'canonical');
            element.setAttribute('href', url);
            document.head.appendChild(element);
        }
    };

    const setStructuredData = (data) => {
        const existingScript = document.querySelector('script[type="application/ld+json"]');
        if (existingScript) {
            existingScript.textContent = JSON.stringify(data);
        } else {
            const script = document.createElement('script');
            script.type = 'application/ld+json';
            script.textContent = JSON.stringify(data);
            document.head.appendChild(script);
        }
    };

    return {
        setMeta,
        setStructuredData
    };
}