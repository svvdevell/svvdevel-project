// generate-sitemap.js - Генерація sitemap.xml для пошукових систем
import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// Отримання списку всіх авто з API
async function getAllCars() {
    try {
        const response = await fetch('https://eleganceauto.od.ua/api/cars-sale');
        const data = await response.json();

        if (data.status === 'success' && Array.isArray(data.data)) {
            return data.data;
        }
    } catch (error) {
        console.error('Помилка отримання списку авто:', error);
    }
    return [];
}

// Генерація XML
async function generateSitemap() {
    const baseUrl = 'https://eleganceauto.od.ua';
    const currentDate = new Date().toISOString().split('T')[0];

    // Статичні сторінки
    const staticPages = [
        { url: '/', priority: '1.0', changefreq: 'daily' },
        { url: '/catalog', priority: '0.9', changefreq: 'daily' },
        { url: '/contact', priority: '0.8', changefreq: 'monthly' }
    ];

    // Отримуємо динамічні сторінки (автомобілі)
    const cars = await getAllCars();
    const carPages = cars.map(car => ({
        url: `/cars/${car.id}`,
        priority: '0.7',
        changefreq: 'weekly',
        lastmod: car.updatedAt || car.createdAt
    }));

    // Об'єднуємо всі сторінки
    const allPages = [...staticPages, ...carPages];

    // Генеруємо XML
    let xml = '<?xml version="1.0" encoding="UTF-8"?>\n';
    xml += '<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">\n';

    for (const page of allPages) {
        xml += '  <url>\n';
        xml += `    <loc>${baseUrl}${page.url}</loc>\n`;
        xml += `    <lastmod>${page.lastmod || currentDate}</lastmod>\n`;
        xml += `    <changefreq>${page.changefreq}</changefreq>\n`;
        xml += `    <priority>${page.priority}</priority>\n`;
        xml += '  </url>\n';
    }

    xml += '</urlset>';

    // Зберігаємо файл
    const sitemapPath = path.join(__dirname, 'dist', 'sitemap.xml');
    fs.writeFileSync(sitemapPath, xml, 'utf-8');

    console.log(`✅ Sitemap згенеровано: ${sitemapPath}`);
    console.log(`📊 Всього URL: ${allPages.length} (${staticPages.length} статичних + ${carPages.length} авто)`);

    return sitemapPath;
}

// Запуск
generateSitemap().catch(console.error);

export default generateSitemap;