// prerender.js - Скрипт для генерації статичних HTML з метатегами
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
    // Базова схема для організації (для всіх сторінок)
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
            "streetAddress": "вул. Прикладна, 1",
            "addressLocality": "Одеса",
            "addressRegion": "Одеська область",
            "postalCode": "65000",
            "addressCountry": "UA"
        },
        "geo": {
            "@type": "GeoCoordinates",
            "latitude": "46.4825",
            "longitude": "30.7233"
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
            "https://www.facebook.com/eleganceauto",
            "https://www.instagram.com/eleganceauto"
        ]
    };

    // Для головної сторінки - додаємо LocalBusiness та Service
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

    // Для каталогу - ItemList
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

    // За замовчуванням - базова схема організації
    return JSON.stringify(organizationSchema, null, 2);
}

// Функція для отримання метаданих машини з API
async function getCarMetadata(carId) {
    try {
        const response = await fetch(`https://eleganceauto.od.ua/api/cars-sale/${carId}`, {
            signal: AbortSignal.timeout(5000) // Таймаут 5 секунд
        });
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const data = await response.json();

        if (data.status === 'success' && data.data) {
            const car = data.data;
            const title = `${car.brand} ${car.model} ${car.year} - ${car.price.toLocaleString('uk-UA')} грн | Elegance Auto Одеса`;
            const description = `${car.brand} ${car.model} ${car.year} року, ${car.color}, ${car.fuel}, ${car.transmission}, пробіг ${car.mileage.toLocaleString('uk-UA')} км. Ціна: ${car.price.toLocaleString('uk-UA')} грн. Викуп та продаж авто в Одесі.`.slice(0, 160);

            return {
                path: `/cars/${carId}`,
                title: title,
                description: description,
                keywords: `${car.brand} ${car.model} купити, ${car.brand} ${car.model} Одеса, ${car.brand} ціна, авто ${car.year}, купить ${car.brand} Одесса`,
                ogImage: car.images && car.images[0] ? `https://eleganceauto.od.ua${car.images[0].fileUrl}` : 'https://eleganceauto.od.ua/images/og-default.jpg',
                carData: car
            };
        }
    } catch (error) {
        console.error(`⚠️  Помилка отримання даних для авто ${carId}:`, error.message);
    }
    return null;
}

// Функція для отримання списку всіх авто з API
async function getAllCarIds() {
    try {
        const response = await fetch('https://eleganceauto.od.ua/api/cars-sale', {
            signal: AbortSignal.timeout(10000) // Таймаут 10 секунд
        });
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const data = await response.json();

        if (data.status === 'success' && Array.isArray(data.data)) {
            return data.data.map(car => car.id);
        }
    } catch (error) {
        console.warn('⚠️  API недоступний, продовжуємо без динамічних роутів:', error.message);
    }
    return [];
}

// Читаємо manifest.json, щоб дізнатися справжні назви файлів
const manifestPath = path.join(__dirname, 'dist', 'manifest.json');
let manifest = {};

if (fs.existsSync(manifestPath)) {
    try {
        manifest = JSON.parse(fs.readFileSync(manifestPath, 'utf-8'));
        console.log('✅ manifest.json знайдено');
    } catch (error) {
        console.warn('⚠️  Помилка читання manifest.json:', error.message);
    }
} else {
    console.warn('⚠️  manifest.json не знайдено, використовуємо стандартні шляхи');
}

function getAssets() {
    // Якщо manifest існує та має дані
    if (manifest && Object.keys(manifest).length > 0) {
        const mainEntry = manifest['index.html'] || Object.values(manifest)[0];
        if (mainEntry) {
            return {
                js: mainEntry.file ? '/' + mainEntry.file : '/assets/index.js',
                css: mainEntry.css && mainEntry.css.length > 0 ? '/' + mainEntry.css[0] : '/assets/index.css'
            };
        }
    }
    
    // Fallback - шукаємо файли в директорії dist/assets
    const distAssetsPath = path.join(__dirname, 'dist', 'assets');
    let jsFile = '/assets/index.js';
    let cssFile = '/assets/index.css';
    
    try {
        if (fs.existsSync(distAssetsPath)) {
            const files = fs.readdirSync(distAssetsPath);
            const jsFiles = files.filter(f => f.endsWith('.js') && !f.endsWith('.map'));
            const cssFiles = files.filter(f => f.endsWith('.css'));
            
            if (jsFiles.length > 0) jsFile = `/assets/${jsFiles[0]}`;
            if (cssFiles.length > 0) cssFile = `/assets/${cssFiles[0]}`;
        }
    } catch (error) {
        console.warn('⚠️  Помилка пошуку assets:', error.message);
    }
    
    return { js: jsFile, css: cssFile };
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
  <meta property="og:locale:alternate" content="ru_UA">
  <meta property="og:site_name" content="Elegance Auto">
  
  <!-- Twitter -->
  <meta name="twitter:card" content="summary_large_image">
  <meta name="twitter:url" content="https://eleganceauto.od.ua${route.path}">
  <meta name="twitter:title" content="${route.title}">
  <meta name="twitter:description" content="${route.description}">
  <meta name="twitter:image" content="${route.ogImage}">
  
  <!-- Додаткові SEO теги -->
  <meta name="robots" content="index, follow, max-image-preview:large, max-snippet:-1, max-video-preview:-1">
  <meta name="googlebot" content="index, follow">
  <link rel="canonical" href="https://eleganceauto.od.ua${route.path}">
  
  <!-- Geo теги для локального SEO -->
  <meta name="geo.region" content="UA-51">
  <meta name="geo.placename" content="Одеса">
  <meta name="geo.position" content="46.4825;30.7233">
  <meta name="ICBM" content="46.4825, 30.7233">
  
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

    console.log('🚀 Початок prerendering з structured data...\n');

    // Перевіряємо чи існує dist директорія
    if (!fs.existsSync(distPath)) {
        console.error('❌ Директорія dist/ не знайдена! Спочатку виконайте build.');
        process.exit(1);
    }

    // Додаємо динамічні роути для машин
    const carIds = await getAllCarIds();
    
    if (carIds.length > 0) {
        console.log(`📦 Знайдено ${carIds.length} автомобілів для prerendering\n`);
        
        for (const carId of carIds) {
            const carRoute = await getCarMetadata(carId);
            if (carRoute) {
                routes.push(carRoute);
            }
        }
    } else {
        console.log('📦 Динамічні роути недоступні, використовуємо тільки статичні сторінки\n');
    }

    // Генеруємо HTML для кожного роуту
    let successCount = 0;
    let errorCount = 0;
    
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

            console.log(`✅ ${route.path}`);
            successCount++;
        } catch (error) {
            console.error(`❌ Помилка для ${route.path}:`, error.message);
            errorCount++;
        }
    }

    console.log(`\n🎉 Prerendering завершено!`);
    console.log(`   ✅ Успішно: ${successCount} сторінок`);
    if (errorCount > 0) {
        console.log(`   ❌ Помилок: ${errorCount}`);
    }
}

// Запуск
prerender().catch(error => {
    console.error('💥 Критична помилка:', error);
    process.exit(1);
});