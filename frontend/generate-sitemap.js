import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

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

async function generateSitemap() {
    const baseUrl = 'https://eleganceauto.od.ua';
    const currentDate = new Date().toISOString().split('T')[0];

    const staticPages = [
        { url: '/', priority: '1.0', changefreq: 'daily', lastmod: currentDate },
        { url: '/catalog', priority: '0.9', changefreq: 'daily', lastmod: currentDate },
        { url: '/contact', priority: '0.8', changefreq: 'monthly', lastmod: currentDate }
    ];

    const cars = await getAllCars();

    const carPages = cars.map(car => {
        const lastmod = (car.updatedAt || car.createdAt || new Date()).split('T')[0];
        return {
            url: `/cars/${car.id}`,
            priority: '0.7',
            changefreq: 'weekly',
            lastmod
        };
    });

    const allPages = [...staticPages, ...carPages];

    let xml = '<?xml version="1.0" encoding="UTF-8"?>\n';
    xml += '<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">\n';

    for (const page of allPages) {
        xml += '  <url>\n';
        xml += `    <loc>${baseUrl}${page.url}</loc>\n`;
        xml += `    <lastmod>${page.lastmod}</lastmod>\n`;
        xml += `    <changefreq>${page.changefreq}</changefreq>\n`;
        xml += `    <priority>${page.priority}</priority>\n`;
        xml += '  </url>\n';
    }

    xml += '</urlset>';

    const distDir = path.join(__dirname, 'dist');
    if (!fs.existsSync(distDir)) fs.mkdirSync(distDir, { recursive: true });

    const sitemapPath = path.join(distDir, 'sitemap.xml');
    fs.writeFileSync(sitemapPath, xml, 'utf-8');

    console.log(`✅ Sitemap згенеровано: ${sitemapPath}`);
    console.log(`📊 Всього URL: ${allPages.length}`);
}

generateSitemap().catch(console.error);
