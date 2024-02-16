import path from 'path'
import webpack from "webpack";
import HtmlWebpackPlugin from "html-webpack-plugin";
import type {Configuration as DevServerConfiguration} from "webpack-dev-server";
import MiniCssExtractPlugin from "mini-css-extract-plugin";
import Dotenv  from "dotenv-webpack";

type Mode = 'production' | 'development'

// пример передачи переменных окружения в dev-server
// npm run start -- --env port=7000
interface EnvVariables {
    mode: Mode,
    port: number,
}

export default (env: EnvVariables) => {

    const isDev = env.mode === 'development';
    const isProd = env.mode === 'production';

    const config: webpack.Configuration = {
        // режим сборки
        mode: env.mode ?? 'development',
        // указываем точку входа, откуда будет начинаться сборка
        entry: path.resolve(__dirname, 'src', 'index.tsx'),
        output: {
            // дирректория сборки
            path: path.resolve(__dirname, 'build'),
            // указываем что в названии файла будет его хэш. Делается для того, чтобы при каждом изменении скриптов,
            // создавалось новое название файл, т.к. браузер кэширует файлы по названию
            filename: '[name].[contenthash].js',
            // при каждой сборке, очищать дирректорию сборки
            clean: true,
        },
        plugins: [
            // указываем на основе какого шаблона будет создаваться html-файл, в который будут подключаться скрипты
            new HtmlWebpackPlugin({template: path.resolve(__dirname, 'public', 'index.html')}),
            // показывает процент выполнения сборки
            isDev && new webpack.ProgressPlugin(),
            // позволяет использовать сss-loader, который минимизирует css в отдельные файлы
            isProd && new MiniCssExtractPlugin({
                filename: 'css/[name].[contenthash:8].css',
                chunkFilename: 'css/[name].[contenthash:8].css',
            }),
            new Dotenv(),
        ].filter(Boolean),
        module: {
            rules: [
                {
                    test: /\.css$/i,
                    // порядок следования лоудеров имеет значение
                    use: [
                        isDev ? "style-loader" : MiniCssExtractPlugin.loader,
                        "css-loader"
                    ],
                },
                {
                    test: /\.s[ac]ss$/i,
                    use: [
                        // Creates `style` nodes from JS strings
                        isDev ? "style-loader" : MiniCssExtractPlugin.loader,
                        // Translates CSS into CommonJS
                        "css-loader",
                        // Compiles Sass to CSS
                        "sass-loader",
                    ],
                },
                {
                    // указываем расширения файлов(.ts и .tsx), которые будут обрабатываться лоудером
                    test: /\.tsx?$/,
                    // название лоудера
                    use: 'ts-loader',
                    // указываем дирректории, которые лоудер не будет обрабатывать
                    exclude: /node_modules/,
                },
            ],
        },
        resolve: {
            // указываем расширения файлов(с исходным кодом), при импорте которых, можно будет не указывать их расширения
            extensions: ['.tsx', '.ts', '.js'],
        },
        devtool: isDev && 'inline-source-map',
        devServer: isDev ? {
            port: env.port ?? 3000,
            open: true,
        } : undefined,
    }
    return config
}
