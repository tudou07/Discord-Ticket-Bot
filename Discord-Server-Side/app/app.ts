import express from 'express';
import bodyParser from 'body-parser';
import cors from 'cors';

import { Routes } from './server/routes';
import { DiscordBot } from './server/discordBot';

class App {
    public app: express.Application;
    public routePrv: Routes = new Routes();
    public discordBot: DiscordBot = new DiscordBot();
    
    constructor() {
        this.app = express();
        this.routePrv.routes(this.app);
        this.config();
    }

    private config(): void {
        this.app.use((req, res, next) => {
            res.header("Access-Control-Allow-Origin", "*");
            res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
            res.setHeader('Access-Control-Allow-Methods', 'GET, POST, OPTIONS, PUT, PATCH, DELETE');
            next();
        })

        // support application/json type post data
        this.app.use(bodyParser.json());
        //support application/x-www-form-urlencoded post data
        this.app.use(bodyParser.urlencoded({ extended: false }));
        this.app.use(cors());
    }
}

export default new App().app;