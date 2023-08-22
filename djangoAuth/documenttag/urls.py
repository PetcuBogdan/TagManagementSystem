from django.urls import path
from documenttag import views
from documenttag.views import Search


urlpatterns = [
    path('searchdocs/', views.document_list),
    path('search/<int:id>', views.document_detail),
    path('search/', views.Search.as_view(), name='search'),
    ]