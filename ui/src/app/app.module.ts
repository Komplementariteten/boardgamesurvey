import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BbgsearchComponent } from './bbgsearch/bbgsearch.component';
import { FormsModule } from '@angular/forms';
import { GameComponent } from './game/game.component';
import { SurveyComponent } from './survey/survey.component';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { VoteitemComponent } from './voteitem/voteitem.component'
import {MatIconModule} from '@angular/material/icon';
import {MatTabsModule} from '@angular/material/tabs';
import { CategorieComponent } from './categorie/categorie.component';
import { DelayedInputModule } from './delayed-input/delayed-input.module';
import { InfotextComponent } from './infotext/infotext.component';



@NgModule({
  declarations: [
    AppComponent,
    BbgsearchComponent,
    GameComponent,
    SurveyComponent,
    VoteitemComponent,
    CategorieComponent,
    InfotextComponent
  ],
  imports: [
    BrowserModule,
    MatTabsModule,
    HttpClientModule,
    FormsModule,
    AppRoutingModule,
    DelayedInputModule,
    MatIconModule,
    NoopAnimationsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
