from re import template
from django.urls import path, include, re_path
from django.views.generic import TemplateView
from documenttag import views


urlpatterns = [
    path('documents/', views.document_list),
    path('documents/<int:id>', views.document_detail),
    path('search/', views.search_document),
]

