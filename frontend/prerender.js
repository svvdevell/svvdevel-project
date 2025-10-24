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
        title: 'Elegance Auto - Викуп та продаж автомобілів в Одесі',
        description: 'Вигідний викуп вашого автомобіля та продаж якісних авто в Одесі. Швидка оцінка, чесні умови, широкий вибір автомобілів.',
        keywords: 'викуп авто Одеса, продаж автомобілів Одеса, автосалон Одеса, купити авто Одеса',
        ogImage: 'https://eleganceauto.od.ua/images/og-home.jpg'
    },
    {
        path: '/contact',
        title: 'Контакти - Elegance Auto | Викуп автомобілів в Одесі',
        description: 'Зв\'яжіться з нами для викупу вашого автомобіля. Працюємо в Одесі та області. Швидка оцінка та вигідні умови.',
        keywords: 'контакти автосалон Одеса, викуп авто телефон, адреса автосалону',
        ogImage: 'https://eleganceauto.od.ua/images/og-contact.jpg'
    },
    {
        path: '/catalog',
        title: 'Каталог автомобілів - Elegance Auto Одеса',
        description: 'Великий вибір якісних автомобілів в Одесі. Перевірені авто з повною історією, чесні ціни та гарантія.',
        keywords: 'каталог авто Одеса, купити авто Одеса, автомобілі в наявності',
        ogImage: 'https://eleganceauto.od.ua/images/og-catalog.jpg'
    }
];

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
                ogImage: car.images && car.images[0] ? `https://eleganceauto.od.ua${car.images[0].fileUrl}` : 'https://eleganceauto.od.ua/images/og-default.jpg'
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

// Генерація HTML з метатегами
function generateHTML(route) {
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
  
  <!-- Structured Data (JSON-LD) -->
  <script type="application/ld+json">
  {
    "@context": "https://schema.org",
    "@type": "${route.path.includes('/cars/') ? 'Car' : 'AutoDealer'}",
    "name": "${route.title}",
    "description": "${route.description}",
    "url": "https://eleganceauto.od.ua${route.path}",
    "image": "${route.ogImage}"
  }
  </script>
  
  <script type="module" crossorigin src="/assets/index.js"></script>
  <link rel="stylesheet" href="/assets/index.css">
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