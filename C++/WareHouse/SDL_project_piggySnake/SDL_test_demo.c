#define _CRT_SECURE_NO_WARNINGS

#include<SDL.h>
#include<SDL_ttf.h>
#include<SDL_mixer.h>
#include<SDL_image.h>
#include<stdbool.h>
#include<stdio.h>
#include<stdlib.h>
#include<time.h>

#define MAXLEN 500
//���ڿ��߶�
const int Screen_width = 640;
const int Screen_height = 480;

/*
	1.����Ŀ¼��ͷ�ļ�
	2.��Ŀ¼��  lib
	3.������->����->������������е�lib���ּӽ���
	4.DLL path·������ <- ��������
	5.����Ϊģ��
*/

//̰����
typedef struct Snake {
	SDL_Point pos[MAXLEN];//�ߵ�������������
	int size;//��Ⱥ͸߶�
	int len;//�ߵ�ǰ����
	int dir;//�ߵ��ƶ�����
}Snake;
Snake snake;

//����
typedef struct Fruit {
	SDL_Point pos_f;
}Fruit;

Fruit fruit;

void fruit_a(SDL_Point  point , SDL_Renderer* render) {
	SDL_SetRenderDrawColor(render, 0, 0, 255, 255);
	
	SDL_Rect rect_f = { point.x,point.y,10,10 };
	SDL_RenderDrawRect(render, &rect_f);
}


void snake_init(Snake* snake,int len) {
	snake->len = len;
	snake->size = 10;
	snake->dir = SDLK_RIGHT;
	for (int i = 0; i < snake->len; i++) {
		snake->pos[i].x = (snake->len - i - 1) * snake->size;//Ҳ�ɸ���ѭ������
		snake->pos[i].y = 0;
	}

}
void snake_draw(Snake* snake,SDL_Renderer* render) {
	
	SDL_SetRenderDrawColor(render, 255, 0, 0, 255);
	for (int i = 0; i < snake->len; i++) {
		//���ƾ���
		SDL_Rect rect = { snake->pos[i].x,snake->pos[i].y,snake->size,snake->size};
		SDL_RenderDrawRect(render, &rect);
	}
}
//�����귢���仯����
void snake_move(Snake* snake) {
	//ֻҪ����ͷ�伴�ɣ������Ǹ�����ͷ�ߵ�
	for (int i = snake->len - 1; i > 0; i--) //��bug
	{
		snake->pos[i] = snake->pos[i - 1];
	}
	//���߸ı䷽��
	switch (snake->dir) {
	case SDLK_UP: //�����·��Դ�����
		snake->pos[0].y -= snake->size;
		break;
	case SDLK_DOWN:
		snake->pos[0].y += snake->size;
		break;
	case SDLK_LEFT:
		snake->pos[0].x -= snake->size;
		break;
	case SDLK_RIGHT:
		snake->pos[0].x += snake->size;
		break;
	}
}

SDL_Point random(SDL_Point point){
	point.x = 0;
	point.y= 0;
	int m, n;
	time_t t;
	srand((unsigned)time(&t));

	m = rand() % 62 + 1;
	n = rand() % 46 + 1;
	point.x = m * 10;
	point.y = n * 10;
    return point;
}

void draw(SDL_Renderer* render) {
	//���ô��ڱ�����ɫ
	SDL_SetRenderDrawColor(render, 255, 255, 255, 255);//RGBAģʽ    �� 0~255 0:��ȫ͸��
	//�õ�ǰ��ɫ���������Ļ
	SDL_RenderClear(render);

	SDL_SetRenderDrawColor(render, 200, 200, 200, 255);

	//������
	for (size_t i = 0; i < Screen_height; i++) {
		SDL_RenderDrawLine(render, 0, i * 10, Screen_width, i * 10);
	}
	for (size_t i = 0; i < Screen_width; i++) {
		SDL_RenderDrawLine(render, i * 10, 0, i * 10, Screen_height);
	}
	//������

	snake_draw(&snake, render);

	//����Ⱦ������ȫ�����Ƶ�������
	SDL_RenderPresent(render);
}

void drawFruit(SDL_Renderer* render) {
	fruit_a(fruit.pos_f, render);

	//����Ⱦ������ȫ�����Ƶ�������
	SDL_RenderPresent(render);
}

void init_sdl() 
{
	snake_init(&snake,4);
}
//void studyDataType() {
//	SDL_Point point;
//	SDL_Rect rect;
//	SDL_Color color;//RGBA ����alphaͨ�� ���� ͸����
//
//	SDL_Window* window = NULL;
//	SDL_Renderer* render = NULL;
//	SDL_Surface* Surface = NULL;
//	SDL_Texture* texture = NULL;
//	
//
//	SDL_Event msg;
//	//�洢����x,y����
//	msg.motion.x;
//	msg.motion.y; 
//
//	//�洢���̰��µļ�
//	msg.key.keysym.sym;
//	
//	//�洢��갴�����ĸ���
//	msg.button.button;
//
//}
//void createWindow(int width,int height) {
//	SDL_Window* w = SDL_CreateWindow("cool", 100, 100, width, height, 4);
//}

Uint32* snakeMoveCall(Uint32 interval, void* param) {
	snake_move(&snake);
	return interval;
}

//������̰����¼�
bool keyPressEventt(SDL_KeyboardEvent* ev) {
	if(ev->keysym.sym == SDLK_SPACE)
	{ 	return true;
	}
	return false;
}


//������̰����¼�
void keyPressEvent(SDL_KeyboardEvent * ev) {
	switch (ev->keysym.sym)
	{
	case SDLK_UP: //�����·��Դ�����
	/*	SDL_Log("up\n");*/
		if (snake.dir == SDLK_DOWN)
			break;
		snake.dir = SDLK_UP;
		break;
	case SDLK_DOWN:
	/*	SDL_Log("Down\n");*/
		if (snake.dir == SDLK_UP) {
			snake.dir == SDLK_UP;
			break;
		}
		snake.dir = SDLK_DOWN;
		break;
	case SDLK_LEFT:
	/*	SDL_Log("Left\n");*/
		if (snake.dir == SDLK_RIGHT) {
			snake.dir == SDLK_RIGHT;
				break;
		}
		snake.dir = SDLK_LEFT;
		break;
	case SDLK_RIGHT:
	/*	SDL_Log("Right\n");*/
		if (snake.dir == SDLK_LEFT) {
			snake.dir == SDLK_LEFT;
				break;
		}
		snake.dir = SDLK_RIGHT;
		break;
	}
}
int showone() {
	// ��ʾһ������������ť����Ϣ��
	SDL_MessageBoxButtonData buttons[] = {
		{ SDL_MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT, 1, "Quit" },
		{ SDL_MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT, 2, "Retry" }
	};

	SDL_MessageBoxData messageboxdata = {
		SDL_MESSAGEBOX_INFORMATION,  // ����
		NULL,                        // ����
		"\xe6\x8f\x90\xe7\xa4\xba",                      // ����
		"\xe8\xaf\xb7\xe9\x80\x89\xe6\x8b\xa9\xe6\x93\x8d\xe4\xbd\x9c",                 // ����
		SDL_arraysize(buttons),       // ��ť����
		buttons,                     // ��ť����
		NULL                          // Ĭ�ϰ�ť
	};

	int buttonid;
	SDL_ShowMessageBox(&messageboxdata, &buttonid);

	// �����û�ѡ��İ�ťִ����Ӧ����
	return buttonid;
}

int main(int argc, char* argv[])
{
	//SDL_Init(SDL_EVERYTHING)
	if (SDL_Init(SDL_INIT_VIDEO|SDL_INIT_TIMER|SDL_INIT_AUDIO)) {
		SDL_Log("%s", SDL_GetError());
		return -1;
	}
	TTF_Init();
	IMG_Init(IMG_INIT_JPG);

	Mix_OpenAudio(44100, MIX_DEFAULT_FORMAT, 2, 2048);
	//�������ں���Ⱦ��(һ���ڴ棬�ݴ��ͼ��Ϣ��˫����)
	SDL_Window* window = NULL;//����ָ��
	SDL_Renderer* render = NULL;//��Ⱦ��ָ��
	SDL_CreateWindowAndRenderer(Screen_width, Screen_height, 4, &window,&render);
	SDL_SetWindowTitle(window,"\xe8\xb4\xaa\xe5\x90\x83\xe8\x9b\x87(\xe7\xbb\x8f\xe5\x85\xb8\xe6\xa8\xa1\xe5\xbc\x8f)");//ʹ��python�� encode����ת��Ϊutf8��ʽ
	SDL_Surface* icon = SDL_LoadBMP("xiaohui.bmp");
	SDL_SetWindowIcon(window, icon);
	
	SDL_Surface* imageSurface = IMG_Load("background.bmp");
	SDL_Texture* imageTexture = SDL_CreateTextureFromSurface(render, imageSurface);
	SDL_FreeSurface(imageSurface);


	SDL_Event event;
	bool isdonee = true;
	while (isdonee) {
		//	SDL_RenderCopy(render, imageTexture, NULL, NULL);
		SDL_RenderPresent(render);
		//
		while (SDL_PollEvent(&event)) {
			if (event.type == SDL_MOUSEBUTTONDOWN || (event.type == SDL_KEYDOWN && keyPressEventt(&event.key))) {
				// �����Ⱦ��
				SDL_RenderClear(render);

				// ���Ƶڶ������������
				isdonee = false;
				draw(render);
				drawFruit(render);

				SDL_RenderPresent(render);
			}
		}
	}

	SDL_AddTimer(200, snakeMoveCall, NULL); // ��ʱ��
a:
	init_sdl();

	fruit.pos_f = random(fruit.pos_f);
	
	TTF_Font* font = TTF_OpenFont("ERASMD.ttf", 24);
	SDL_Color textColor = { 0,0,0 };
	
	// ��Ⱦ����������
	SDL_Texture* scoreTexture = NULL;


	Mix_Music* bgm = Mix_LoadMUS("background_music.mp3");  // ���ر��������ļ�
	if (bgm == NULL) {
		return 1;
	}
	
	Mix_Chunk* soundEffect = Mix_LoadWAV("sound_effect.wav");
	if (soundEffect == NULL) {
		// �������ʧ�ܵ����
		return 1;
	}
	// ���ű�������
	Mix_PlayMusic(bgm, -1);  // -1��ʾѭ������
	//����һ����ʱ�� ԭλ��
	//�¼�����



	bool isdone = false;
	while (!isdone) {
		//֡��   fps ���24֡  һ��60֡   --ÿ���ͼ�Ĵ���
		//��ȡ��ǰʱ��
		Uint32 begTime = SDL_GetTicks(); //��ȡ���ں�����

		draw(render);
		drawFruit(render);

		//�����¼�����
		SDL_Event ev;

		//��ȡ�¼�,��һ��ѭ���� �������¼����������
		while (SDL_PollEvent(&ev)) {
			switch (ev.type) {
			case SDL_QUIT:  //�رմ����¼�
				isdone = true;
				break;
			case SDL_KEYDOWN://��������
				keyPressEvent(&ev.key);
				break;
			}
		}
		char scoreText[50];

		sprintf(scoreText, "Score: %d", snake.len - 4);
		if (scoreTexture) {
			SDL_DestroyTexture(scoreTexture);
		}
		SDL_Surface* surface_1 = TTF_RenderText_Solid(font, scoreText, textColor);

		scoreTexture = SDL_CreateTextureFromSurface(render, surface_1);
		SDL_Rect textRect = { 10, 10, surface_1->w, surface_1->h };

		SDL_RenderCopy(render, scoreTexture, NULL, &textRect);
		// ���´���
		SDL_RenderPresent(render);
		// �ͷű��������
		SDL_FreeSurface(surface_1);
		if (scoreTexture) {
			SDL_DestroyTexture(scoreTexture);
		}
		if (fruit.pos_f.x == snake.pos[0].x && fruit.pos_f.y == snake.pos[0].y)
		{
			SDL_Log("%d,%d,%d", snake.pos[0].x, snake.pos[0].y, snake.len + 1);
			fruit.pos_f = random(fruit.pos_f);
			snake.len++;
		}

		if ((snake.pos[0].x == 630 && snake.dir == SDLK_RIGHT) || (snake.pos[0].y == 470 && snake.dir == SDLK_DOWN) || (snake.pos[0].x == 0 && snake.dir == SDLK_LEFT) || (snake.pos[0].y == 0 && snake.dir == SDLK_UP)) {
			Mix_FreeMusic(bgm);  // �ͷű���������Դ
			Mix_PlayChannel(-1, soundEffect, 0);  // -1��ʾʹ�õ�һ�����õ�ͨ����0��ʾ����һ��
			switch (showone()) {
			case 1:  // Quit
				SDL_Quit();
				break;
			case 2:  // Retry
				goto a;
				break;
			default:
				// ��������
				break;
			}
		}
		if (snake.len > 4) {
			for (int i = 1; i < snake.len; i++) {
				if (snake.pos[0].x == snake.pos[i].x && snake.pos[0].y == snake.pos[i].y) {
					Mix_FreeMusic(bgm);  // �ͷű���������Դ
					Mix_PlayChannel(-1, soundEffect, 0);  // -1��ʾʹ�õ�һ�����õ�ͨ����0��ʾ����һ��
					/*if (SDL_ShowSimpleMessageBox(SDL_MESSAGEBOX_INFORMATION, "bad news !", "you lose", window)) {
						SDL_Log("%s", SDL_GetError());
						return -1;
					}*/

					switch (showone()) {
					case 1:  // Quit
						SDL_Quit();
						break;
					case 2:  // Retry
						goto a;
						break;
					default:
						// ��������
						break;
					}
				}
			}
		}
		//����֡��
		Uint32 eaplsed = SDL_GetTicks() - begTime; //����ʱ��
		int delay = 1000 / 60.0 - eaplsed;
		if (delay > 0) SDL_Delay(delay);//��̬�ӳ�
	}

	Mix_FreeMusic(bgm);  // �ͷű���������Դ
	Mix_FreeChunk(soundEffect);
	Mix_CloseAudio();  // �ر���Ƶ�豸
	//��̬�ڴ���ͷ�
	SDL_FreeSurface(icon);
	TTF_CloseFont(font);
	SDL_DestroyWindow(window);
	SDL_DestroyRenderer(render);
	TTF_Quit();
	SDL_Quit();
	return 0;
}