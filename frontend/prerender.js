// prerender.js - Скрипт для генерації статичних HTML з метатегами
import puppeteer from 'puppeteer';
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
function generateStructuredData(route, carData = null) {
    // Базова схема для організації
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

    // Для головної сторінки
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
                        "name": "Elegance Auto"
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

    // Для сторінки контактів
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

    // Для каталогу
    if (route.path === '/catalog') {
        return JSON.stringify({
            "@context": "https://schema.org",
            "@type": "CollectionPage",
            "name": "Каталог автомобілів",
            "url": "https://eleganceauto.od.ua/catalog",
            "provider": organizationSchema
        }, null, 2);
    }

    // Для конкретного автомобіля
    if (carData) {
        return JSON.stringify({
            "@context": "https://schema.org",
            "@graph": [
                organizationSchema,
                {
                    "@type": "Car",
                    "name": `${carData.brand} ${carData.model}`,
                    "brand": {
                        "@type": "Brand",
                        "name": carData.brand
                    },
                    "model": carData.model,
                    "vehicleModelDate": carData.year.toString(),
                    "productionDate": carData.year.toString(),
                    "mileageFromOdometer": {
                        "@type": "QuantitativeValue",
                        "value": carData.mileage,
                        "unitCode": "KMT"
                    },
                    "fuelType": carData.fuel,
                    "vehicleTransmission": carData.transmission,
                    "color": carData.color,
                    "bodyType": carData.bodyType || "Sedan",
                    "vehicleEngine": {
                        "@type": "EngineSpecification",
                        "fuelType": carData.fuel
                    },
                    "image": carData.images && carData.images.length > 0 
                        ? carData.images.map(img => `https://eleganceauto.od.ua${img.fileUrl}`)
                        : ["https://eleganceauto.od.ua/images/og-default.jpg"],
                    "offers": {
                        "@type": "Offer",
                        "price": carData.price,
                        "priceCurrency": "UAH",
                        "availability": "https://schema.org/InStock",
                        "url": `https://eleganceauto.od.ua/cars/${carData.id}`,
                        "seller": {
                            "@type": "AutoDealer",
                            "name": "Elegance Auto"
                        },
                        "itemCondition": "https://schema.org/UsedCondition"
                    },
                    "description": carData.description || `${carData.brand} ${carData.model} ${carData.year} року`,
                    "url": `https://eleganceauto.od.ua/cars/${carData.id}`
                }
            ]
        }, null, 2);
    }

    // За замовчуванням
    return JSON.stringify(organizationSchema, null, 2);
}

// Функція для отримання метаданих машини з API
async function getCarMetadata(carId) {
    try {
        const response = await fetch(`https://eleganceauto.od.ua/api/cars-sale/${carId}`);
        const data = await response.json();

        if (data.status === 'success' && data.data) {
            const car = data.data;
            const title = `${car.brand} ${car.model} ${car.year} - ${car.price.toLocaleString('uk-UA')} грн`;
            const description = `${car.brand} ${car.model} ${car.year} року, ${car.color}, ${car.fuel}, ${car.transmission}, пробіг ${car.mileage.toLocaleString('uk-UA')} км. Ціна: ${car.price.toLocaleString('uk-UA')} грн. ${car.description || ''}`.slice(0, 160);

            return {
                path: `/cars/${carId}`,
                title: title,
                description: description,
                keywords: `${car.brand} ${car.model} купити, ${car.brand} ${car.model} Одеса, ${car.brand} ціна, авто ${car.year}`,
                ogImage: car.images && car.images[0] ? `https://eleganceauto.od.ua${car.images[0].fileUrl}` : 'https://eleganceauto.od.ua/images/og-default.jpg',
                carData: car
            };
        }
    } catch (error) {
        console.error(`Помилка отримання даних для авто ${carId}:`, error);
    }
    return null;
}

// Функція для отримання списку всіх авто з API
async function getAllCarIds() {
    try {
        const response = await fetch('https://eleganceauto.od.ua/api/cars-sale');
        const data = await response.json();

        if (data.status === 'success' && Array.isArray(data.data)) {
            return data.data.map(car => car.id);
        }
    } catch (error) {
        console.error('Помилка отримання списку авто:', error);
    }
    return [];
}

// ВИПРАВЛЕНА ФУНКЦІЯ: Функція для отримання assets без manifest.json
function getAssets() {
    const distAssetsPath = path.join(__dirname, 'dist', 'assets');
    
    try {
        if (fs.existsSync(distAssetsPath)) {
            const files = fs.readdirSync(distAssetsPath);
            
            // Шукаємо JS файл (не .map файли)
            const jsFile = files.find(f => f.startsWith('index-') && f.endsWith('.js') && !f.endsWith('.map'));
            // Шукаємо CSS файл
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
    
    // Fallback - повертаємо дефолтні шляхи
    console.warn('⚠️  Використовуємо дефолтні шляхи для assets');
    return {
        js: '/assets/index.js',
        css: '/assets/index.css'
    };
}

// Генерація HTML з метатегами
function generateHTML(route) {
    const { js, css } = getAssets();
    const structuredData = generateStructuredData(route, route.carData);
    
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
                <link rel="shortcut icon" href="./public/favicon.ico" type="image/x-icon">
                <link rel="stylesheet" href="/fonts/SFpro/stylesheet.css">
                <link rel="manifest" href="/manifest.webmanifest">
                
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

    // Додаємо динамічні роути для машин
    const carIds = await getAllCarIds();
    console.log(`📦 Знайдено ${carIds.length} автомобілів для prerendering\n`);

    for (const carId of carIds) {
        const carRoute = await getCarMetadata(carId);
        if (carRoute) {
            routes.push(carRoute);
        }
    }

    // Генеруємо HTML для кожного роуту
    for (const route of routes) {
        try {
            const html = generateHTML(route);

            // Створюємо директорію для роуту
            const routePath = route.path === '/' ? 'index.html' : route.path;
            const fullPath = path.join(distPath, routePath);
            const dir = path.dirname(fullPath);

            if (!fs.existsSync(dir)) {
                fs.mkdirSync(dir, { recursive: true });
            }

            // Записуємо HTML файл
            const fileName = routePath.endsWith('.html') ? fullPath : path.join(fullPath, 'index.html');
            fs.writeFileSync(fileName, html, 'utf-8');

            console.log(`✅ ${route.path} -> ${fileName}`);
        } catch (error) {
            console.error(`❌ Помилка для ${route.path}:`, error);
        }
    }

    console.log(`\n🎉 Prerendering завершено! Оброблено ${routes.length} сторінок.`);
}

// Запуск
prerender().catch(console.error);