import axios from 'axios';
import OpenAI from 'openai';

require('dotenv').config();

class GPTService {
    private openai: OpenAI;

    constructor() {
        this.openai = new OpenAI({
            apiKey: process.env.OPENAI_API_KEY
        })
    }

    public async getGPTResponse(prompt: string): Promise<boolean> {
        try {
            const response = await this.openai.chat.completions.create({
                model: "gpt-3.5-turbo",
                messages: [
                  {
                    "role": "user",
                    "content": 'Do you think from the below message that user is having an issue? Only answer in true or false' + prompt
                  }
                ],
                temperature: 1,
                max_tokens: 256,
                top_p: 1,
                frequency_penalty: 0,
                presence_penalty: 0,
              });

            const result = response.choices[0].message.content ?? 'false';
            console.log(response.choices[0].message.content);

            if (result.toLowerCase().startsWith('true')) return true;
            else return false;

        } catch (error) {
            console.error(error);
            return error;
        }
    }
}

export default GPTService;