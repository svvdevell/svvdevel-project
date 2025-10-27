// prerender.js - Версія БЕЗ генерації сторінок авто
import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// Конфігурація роутів з SEO метаданими українською
const routes = [
    {
        path: '/',
        title: 'Авто викуп в Одесі - Elegance Auto | Швидкий викуп автомобілів',
        description: 'Терміновий викуп авто в Одесі за вигідною ціною. Оцінка за 15 хвилин, виплата готівкою. Викупаємо авто в будь-якому стані. Выкуп авто Одесса - быстро, выгодно, надежно.',
        keywords: 'авто викуп Одеса, викуп автомобілів Одеса, терміновий викуп авто, выкуп авто Одесса, продать авто Одесса, скупка автомобилей, автовыкуп, викуп авто готівкою, продаж автомобілів Одеса, автосалон Одеса',
        ogImage: 'https://eleganceauto.od.ua/images/og-home.jpg'
    },
    {
        path: '/contact',
        title: 'Контакти - Викуп авто Одеса | Elegance Auto',
        description: 'Терміновий викуп автомобілів в Одесі та області. Телефонуйте зараз для оцінки вашого авто. Виїзд оцінювача безкоштовно. Выкуп авто в Одессе - звоните!',
        keywords: 'викуп авто Одеса контакти, викуп авто телефон Одеса, выкуп авто Одесса телефон, продать машину Одесса, автовыкуп контакты, адреса автосалону Одеса, терміновий викуп авто',
        ogImage: 'https://eleganceauto.od.ua/images/og-contact.jpg'
    },
    {
        path: '/catalog',
        title: 'Каталог автомобілів - Elegance Auto Одеса | Продаж авто',
        description: 'Купити авто в Одесі після викупу. Великий вибір перевірених автомобілів з повною історією. Авто після trade-in за вигідними цінами. Купить авто в Одессе.',
        keywords: 'каталог авто Одеса, купити авто Одеса, продаж авто Одеса, купить машину Одесса, автомобілі в наявності, б/у авто Одеса, авто после выкупа, trade-in Одеса',
        ogImage: 'https://eleganceauto.od.ua/images/og-catalog.jpg'
    },
];

// Функція для генерації JSON-LD structured data
function generateStructuredData(route) {
    const organizationSchema = {
        "@context": "https://schema.org",
        "@type": "AutoDealer",
        "name": "Elegance Auto",
        "description": "Викуп та продаж автомобілів в Одесі",
        "url": "https://eleganceauto.od.ua",
        "logo": "https://eleganceauto.od.ua/images/logo.png",
        "image": "https://eleganceauto.od.ua/images/og-home.jpg",
        "telephone": "+380 (48) 123-45-67",
        "address": {
            "@type": "PostalAddress",
            "streetAddress": "Полковника Гуляєва, 107/1, Лиманка",
            "addressLocality": "Одеса",
            "addressRegion": "Одеська область",
            "postalCode": "65104",
            "addressCountry": "UA"
        },
        "geo": {
            "@type": "GeoCoordinates",
            "latitude": "46.384590958544315",
            "longitude": "30.704781900546564"
        },
        "openingHoursSpecification": [
            {
                "@type": "OpeningHoursSpecification",
                "dayOfWeek": ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday"],
                "opens": "09:00",
                "closes": "18:00"
            },
            {
                "@type": "OpeningHoursSpecification",
                "dayOfWeek": "Saturday",
                "opens": "10:00",
                "closes": "16:00"
            }
        ],
        "priceRange": "$$",
        "sameAs": [
            "https://t.me/eleganceauto_odessa",
            "https://www.instagram.com/elegance_auto_od"
        ]
    };

    if (route.path === '/') {
        return JSON.stringify({
            "@context": "https://schema.org",
            "@graph": [
                organizationSchema,
                {
                    "@type": "Service",
                    "serviceType": "Викуп автомобілів",
                    "provider": {
                        "@type": "AutoDealer",
                        "name": "Elegance Auto",
                        "address": {
                            "@type": "PostalAddress",
                            "streetAddress": "Полковника Гуляєва, 107/1, Лиманка",
                            "addressLocality": "Одеса",
                            "addressRegion": "Одеська область",
                            "postalCode": "65104",
                            "addressCountry": "UA"
                        },
                    },
                    "areaServed": {
                        "@type": "City",
                        "name": "Одеса"
                    },
                    "description": "Терміновий викуп автомобілів в Одесі. Оцінка за 15 хвилин, виплата готівкою.",
                    "offers": {
                        "@type": "Offer",
                        "availability": "https://schema.org/InStock"
                    }
                },
                {
                    "@type": "WebSite",
                    "url": "https://eleganceauto.od.ua",
                    "name": "Elegance Auto",
                    "potentialAction": {
                        "@type": "SearchAction",
                        "target": "https://eleganceauto.od.ua/catalog?search={search_term_string}",
                        "query-input": "required name=search_term_string"
                    }
                }
            ]
        }, null, 2);
    }

    if (route.path === '/contact') {
        return JSON.stringify({
            "@context": "https://schema.org",
            "@graph": [
                organizationSchema,
                {
                    "@type": "ContactPage",
                    "url": "https://eleganceauto.od.ua/contact",
                    "name": "Контакти - Elegance Auto"
                }
            ]
        }, null, 2);
    }

    if (route.path === '/catalog') {
        return JSON.stringify({
            "@context": "https://schema.org",
            "@type": "CollectionPage",
            "name": "Каталог автомобілів",
            "url": "https://eleganceauto.od.ua/catalog",
            "provider": organizationSchema
        }, null, 2);
    }

    return JSON.stringify(organizationSchema, null, 2);
}

// Функція для отримання assets
function getAssets() {
    const distAssetsPath = path.join(__dirname, 'dist', 'assets');
    
    try {
        if (fs.existsSync(distAssetsPath)) {
            const files = fs.readdirSync(distAssetsPath);
            
            const jsFile = files.find(f => f.startsWith('index-') && f.endsWith('.js') && !f.endsWith('.map'));
            const cssFile = files.find(f => f.startsWith('index-') && f.endsWith('.css'));
            
            if (jsFile && cssFile) {
                console.log(`📦 Знайдено assets: JS=${jsFile}, CSS=${cssFile}`);
                return {
                    js: `/assets/${jsFile}`,
                    css: `/assets/${cssFile}`
                };
            } else {
                console.warn(`⚠️  Assets неповні: JS=${jsFile}, CSS=${cssFile}`);
            }
        } else {
            console.warn(`⚠️  Директорія ${distAssetsPath} не існує`);
        }
    } catch (error) {
        console.error('⚠️  Помилка читання assets:', error);
    }
    
    console.warn('⚠️  Використовуємо дефолтні шляхи для assets');
    return {
        js: '/assets/index.js',
        css: '/assets/index.css'
    };
}

// Генерація HTML з метатегами
function generateHTML(route) {
    const { js, css } = getAssets();
    const structuredData = generateStructuredData(route);
    
    return `<!DOCTYPE html>
        <html lang="uk">
                <head>
                <meta charset="UTF-8">
                <meta name="viewport" content="width=device-width, initial-scale=1.0">
                <!-- Основні мета-теги -->
                <title>${route.title}</title>
                <meta name="description" content="${route.description}">
                <meta name="keywords" content="${route.keywords}">
                
                <!-- Open Graph / Facebook -->
                <meta property="og:type" content="website">
                <meta property="og:url" content="https://eleganceauto.od.ua${route.path}">
                <meta property="og:title" content="${route.title}">
                <meta property="og:description" content="${route.description}">
                <meta property="og:image" content="${route.ogImage}">
                <meta property="og:locale" content="uk_UA">
                
                <!-- Twitter -->
                <meta name="twitter:card" content="summary_large_image">
                <meta name="twitter:url" content="https://eleganceauto.od.ua${route.path}">
                <meta name="twitter:title" content="${route.title}">
                <meta name="twitter:description" content="${route.description}">
                <meta name="twitter:image" content="${route.ogImage}">
                
                <!-- Додаткові SEO теги -->
                <meta name="robots" content="index, follow">
                <meta name="googlebot" content="index, follow">

                <link rel="canonical" href="https://eleganceauto.od.ua${route.path}">
                <link rel="stylesheet" href="/fonts/SFpro/stylesheet.css">
                <link rel="manifest" href="/manifest.webmanifest">

                <!-- Favicon -->
                <link rel="icon" type="image/x-icon" href="/favicon.ico">
                <link rel="shortcut icon" type="image/x-icon" href="/favicon.ico">
                <link rel="apple-touch-icon" sizes="180x180" href="/favicon.ico">
                <link rel="icon" type="image/png" sizes="32x32" href="/favicon.ico">
                <link rel="icon" type="image/png" sizes="16x16" href="/favicon.ico">
                
                <!-- Structured Data (JSON-LD) -->
                <script type="application/ld+json">
                    ${structuredData}
                </script>
                
                <script type="module" crossorigin src="${js}"></script>
                <link rel="stylesheet" href="${css}">
            </head>
            <body>
                <div id="app"></div>
            </body>
        </html>`;
}

// Основна функція prerendering
async function prerender() {
    const distPath = path.join(__dirname, 'dist');

    console.log('🚀 Початок prerendering...\n');
    console.log('ℹ️  Генеруємо тільки основні сторінки (/, /contact, /catalog)');
    console.log('ℹ️  Сторінки окремих авто НЕ генеруються (API недоступний під час збірки)\n');

    // Генеруємо HTML для кожного роуту
    for (const route of routes) {
        try {
            const html = generateHTML(route);

            // Правильне створення шляху та директорій
            let fileName;
            
            if (route.path === '/') {
                fileName = path.join(distPath, 'index.html');
            } else {
                const dirPath = path.join(distPath, route.path);
                
                if (!fs.existsSync(dirPath)) {
                    fs.mkdirSync(dirPath, { recursive: true });
                    console.log(`📁 Створено директорію: ${dirPath}`);
                }
                
                fileName = path.join(dirPath, 'index.html');
            }

            fs.writeFileSync(fileName, html, 'utf-8');

            console.log(`✅ ${route.path} -> ${fileName}`);
        } catch (error) {
            console.error(`❌ Помилка для ${route.path}:`, error);
        }
    }

    console.log(`\n🎉 Prerendering завершено! Оброблено ${routes.length} сторінок.`);
    console.log(`\nℹ️  Примітка: Для генерації сторінок окремих автомобілів потрібен окремий процес`);
    console.log(`   після того як backend запущений і авто додані в базу даних.`);
}

// Запуск
prerender().catch(console.error);