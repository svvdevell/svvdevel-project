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
        { url: '/contact', priority: '0.8', changefreq: 'monthly', lastmod: currentDate },
        { url: '/blog', priority: '0.7', changefreq: 'monthly', lastmod: currentDate }
    ];

    const cars = await getAllCars();
    const carPages = cars.map(car => {
        const lastmod = car.updatedAt 
            ? new Date(car.updatedAt).toISOString().split('T')[0]
            : currentDate;
        return {
            url: `/cars/${car.id}`,
            priority: '0.7',
            changefreq: 'weekly',
            lastmod
        };
    });

    const allPages = [...staticPages, ...carPages];

    // КРИТИЧНО: используй template literals правильно
    const xml = `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
${allPages.map(page => `  <url>
    <loc>${baseUrl}${page.url}</loc>
    <lastmod>${page.lastmod}</lastmod>
    <changefreq>${page.changefreq}</changefreq>
    <priority>${page.priority}</priority>
  </url>`).join('\n')}
</urlset>`;

    const distDir = path.join(__dirname, 'dist');
    if (!fs.existsSync(distDir)) fs.mkdirSync(distDir, { recursive: true });

    const sitemapPath = path.join(distDir, 'sitemap.xml');
    fs.writeFileSync(sitemapPath, xml, 'utf-8');

    console.log(`✅ Sitemap згенеровано: ${sitemapPath}`);
    console.log(`📊 Всього URL: ${allPages.length}`);
}

generateSitemap().catch(console.error);
