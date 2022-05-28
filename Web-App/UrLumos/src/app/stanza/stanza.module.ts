import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { StanzaPageRoutingModule } from './stanza-routing.module';

import { StanzaPage } from './stanza.page';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    StanzaPageRoutingModule
  ],
  declarations: [StanzaPage]
})
export class StanzaPageModule {}
