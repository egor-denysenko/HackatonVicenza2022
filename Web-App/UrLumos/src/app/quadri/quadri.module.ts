import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { QuadriPageRoutingModule } from './quadri-routing.module';

import { QuadriPage } from './quadri.page';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    QuadriPageRoutingModule
  ],
  declarations: [QuadriPage]
})
export class QuadriPageModule {}
